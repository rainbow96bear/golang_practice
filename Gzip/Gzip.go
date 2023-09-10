package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func main(){
	sampleData := "Hi This is Rainbowbear World."

	// string을 byte 데이터로 변환
	Data := []byte(sampleData)

	// 데이터 압축
	compresssedData, err := compressData(Data)
	if err != nil{
		fmt.Println("Compress fail", err)
	}
	fmt.Println("압축 된 데이터 :", string(compresssedData))

	// 압축된 데이터 풀기
	decompressedData, err := decompressData(compresssedData)
	if err != nil {
		fmt.Println("decompress fail", err)
		return
	}
	fmt.Println("압축이 풀린 데이터 :", string(decompressedData))

	fmt.Println("원본 데이터 :",sampleData)
}

func compressData(data []byte) ([]byte, error){
	buf := new(bytes.Buffer)
	writer := gzip.NewWriter(buf)
	_, err := writer.Write(data)

	if err != nil {
		return nil, err
	}
	writer.Close()
	return buf.Bytes(), nil
}

func decompressData(data []byte) ([]byte, error){
	r, err := gzip.NewReader(bytes.NewReader(data))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	originData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return originData, nil
}