package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Kreagentle/gRPC/primenumbers/proto"
)

func CalculatePrimes(calcserver pb.CalculatorClient) error {
	log.Println("Function CalculatePrimes is invoked")

	req := &pb.PrimeRequest{
		Number1: 120,
	}

	stream, err := calcserver.Calculate(context.Background(), req)
	if err != nil {
		log.Println("Error while CalculatePrimes is invoked")
		return err
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("Error while CalculatePrimes is invoked")
		}

		log.Printf("Prime: %d", res.Number2)
	}

	return nil
}
