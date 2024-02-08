// main.go
package main

import (
	"context"
	"fmt"
	grpcchat "golang_practice/grpc_chat/protoc"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	grpcchat.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients map[grpcchat.ChatService_BroadcastServer]struct{}
}

func (s *server) SendMessage(ctx context.Context, req *grpcchat.MessageRequest) (*grpcchat.MessageResponse, error) {
    message := fmt.Sprintf("[%s]: %s", req.GetUsername(), req.GetMessage())

    s.mu.Lock()
    defer s.mu.Unlock()

    // 로그로 메시지 확인
    log.Printf("Received message: %s", message)

    // 현재 서버에 연결된 모든 클라이언트에게 메시지 브로드캐스트
    for client := range s.clients {
        if err := client.Send(&grpcchat.MessageResponse{Message: message}); err != nil {
            log.Printf("메시지 브로드캐스트 실패: %v", err)
        }
    }

    return &grpcchat.MessageResponse{Message: message}, nil
}


func (s *server) Broadcast(stream grpcchat.ChatService_BroadcastServer) error {
	s.mu.Lock()
	s.clients[stream] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("리스닝 실패: %v", err)
	}

	srv := grpc.NewServer()
	chatServer := &server{
		clients: make(map[grpcchat.ChatService_BroadcastServer]struct{}),
	}
	grpcchat.RegisterChatServiceServer(srv, chatServer)

	log.Println("gRPC 서버가 :50051 포트에서 실행 중입니다.")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
