package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Kreagentle/gRPC/average/proto"
)

func (s *Server) AvgCalculate(stream pb.Calculator_CalculateServer) error {
	log.Printf("Function AvgCalculate is invoked")
	var sum int32
	var count int

	for {
		rec, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Number2: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("cant receive the message in the server side %v", err)
		}

		fmt.Printf("Receiving request %v\n\n", rec.Number)
		sum += rec.Number
		count++
	}
	return nil
}
