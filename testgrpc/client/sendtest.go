package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func SendTest(client pb.ServerClient) {
	fmt.Println("Send text is invoked")
	result, err := client.Test(context.Background(), &pb.TestRequest{TransactionRequest: "1"})
	if err != nil {
		log.Fatalf("error is %v", err)
	}
	log.Printf("The transaction request is %v", result)
}
