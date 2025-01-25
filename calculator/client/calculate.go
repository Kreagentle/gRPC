package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/calculator/proto"
)

func Calculate(client pb.CalculatorClient) {
	fmt.Println("Calculation is started")
	result, err := client.Calculate(context.Background(), &pb.CalculatorRequest{Number1: 1, Number2: 2})
	if err != nil {
		log.Fatalf("error is %v", err)
	}
	log.Printf("The result sum is %v", result.Number3)
}
