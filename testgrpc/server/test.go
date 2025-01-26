package main

import (
	"context"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func (s *Server) Test(c context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Function Test is invoked %v\n\n", in)
	return &pb.TestResponse{
		TransactionResponse: "The id of transaction is " + in.TransactionRequest,
	}, nil
}
