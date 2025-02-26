package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/average/proto"
)

func AvrCalculate(client pb.CalculatorClient) {
	fmt.Println("Avr calculate is invoked")

	request := []int32{1, 2, 3, 4, 5}

	stream, err := client.Calculate(context.Background())
	if err != nil {
		log.Fatalf("error is %v", err)
	}

	for _, req := range request {
		fmt.Printf("Sending request %v\n\n", req)
		err = stream.Send(&pb.AvgRequest{Number: req})
		if err != nil {
			log.Fatalf("error while writing into the stream %v", err)
		}
	}

	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while wrinting response into the stream %v", err)
	}
	log.Printf("The transaction request is %v", result.Number2)
}
