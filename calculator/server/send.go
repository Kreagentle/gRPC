package main

import (
	"context"
	"log"

	pb "github.com/Kreagentle/gRPC/calculator/proto"
)

func (s *Server) Calculate(c context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculation that happend from the server side %v\n", in)
	return &pb.CalculatorResponse{
		Number3: in.Number1 + in.Number2,
	}, nil
}
