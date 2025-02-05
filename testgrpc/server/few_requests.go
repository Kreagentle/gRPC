package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func (s *Server) TestFewRequests(stream pb.Server_TestFewRequestsServer) error {
	log.Printf("Function TestFewRequests is invoked")
	result := ""
	for {
		rec, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.TestResponse{TransactionResponse: result})
		}

		if err != nil {
			log.Fatalf("cant receive the message in the server side %v", err)
		}

		result += fmt.Sprintf("transaction id %s", rec.TransactionRequest)
	}
	return nil
}
