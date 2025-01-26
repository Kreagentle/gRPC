package main

import (
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func (s *Server) TestFewTimes(in *pb.TestRequest, stream pb.Server_TestFewTimesServer) error {
	log.Printf("Function TestFewTimes is invoked %v\n\n", in)
	for i := range []int{1, 2, 3, 4, 5} {
		data := fmt.Sprintf("Invokes %s the data %d", in.TransactionRequest, i)

		stream.Send(&pb.TestResponse{
			TransactionResponse: data,
		})
	}
	return nil
}
