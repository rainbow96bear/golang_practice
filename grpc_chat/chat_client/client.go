package main

import (
	"bufio"
	"context"
	"fmt"
	grpcchat "golang_practice/grpc_chat/protoc"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// gRPC 서버에 연결
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("연결 실패: %v", err)
	}
	defer conn.Close()

	client := grpcchat.NewChatServiceClient(conn)

	// 채팅 사용자 정보 설정
	username := "User1"
	fmt.Fscanf(reader, "%s\n", &username)

	// 채팅 서버로 메시지 전송
	go func() {
		for {
			message, _, _ :=reader.ReadLine()
			sendMessage(client, username, string(message))
		}
	}()

	// 채팅 서버로부터 메시지 수신
	receiveMessages(client, username)
}

func sendMessage(client grpcchat.ChatServiceClient, username, message string) {
	// 채팅 메시지 생성
	req := &grpcchat.MessageRequest{
		Username: username,
		Message:  message,
	}

	// 채팅 서버에 메시지 전송
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("메시지 전송 실패: %v", err)
	}

	fmt.Printf("서버 응답: %s\n", resp.GetMessage())
}

func receiveMessages(client grpcchat.ChatServiceClient, username string) {
	// 채팅 서버로부터 메시지 수신 스트림 설정
	stream, err := client.Broadcast(context.Background())
	if err != nil {
		log.Fatalf("메시지 수신 스트림 설정 실패: %v", err)
	}

	// 채팅 서버로부터 메시지 수신 및 출력
	for {
		req := &grpcchat.MessageRequest{Username: username}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("메시지 전송 실패: %v", err)
		}

		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("메시지 수신 실패: %v", err)
		}

		fmt.Printf("새로운 메시지: %s\n", resp.GetMessage())
	}
}
