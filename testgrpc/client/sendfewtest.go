package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func SendFewTest(client pb.ServerClient) {
	fmt.Println("Send few text is invoked")

	request := &pb.TestRequest{TransactionRequest: "1"}

	stream, err := client.TestFewTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("error is %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error while reading the stream %v", err)
		}

		log.Printf("The transaction request is %v", message.TransactionResponse)
	}
}
