package main

import (
	"context"
	"log"

	pb "github.com/Kreagentle/gRPC/primenumbers/proto"
)

func (s *Server) Calculate(c context.Context, in *pb.PrimeRequest) (*pb.PrimeResponse, error) {
	log.Printf("Calculation that happend from the server side %v\n", in)
	return &pb.PrimeResponse{
		Number3: in.Number1 + in.Number2,
	}, nil
}
