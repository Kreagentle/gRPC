package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func SendFewRequests(client pb.ServerClient) {
	fmt.Println("Send few requests is invoked")

	request := []*pb.TestRequest{{TransactionRequest: "1"}, {TransactionRequest: "2"}, {TransactionRequest: "3"}}

	stream, err := client.TestFewRequests(context.Background())
	if err != nil {
		log.Fatalf("error is %v", err)
	}

	for _, req := range request {
		fmt.Printf("Sending request %v\n\n", req)
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("error while writing into the stream %v", err)
		}
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while wrinting response into the stream %v", err)
	}
	log.Printf("The transaction request is %v", result.TransactionResponse)
}
