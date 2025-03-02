package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

func (s *Server) BiDirectionalRequests(stream pb.Server_TestBiDirectionalServer) error {
	log.Printf("Function BiDirectionalRequests is invoked")
	result := ""
	for {
		rec, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("cant receive the message in the server side %v", err)
		}

		result += fmt.Sprintf("transaction id %s", rec.TransactionRequest)

		err = stream.Send(&pb.TestResponse{
			TransactionResponse: result,
		})
		if err != nil {
			log.Fatalf("cant send the message in the server side %v", err)
		}
	}
	return nil
}
