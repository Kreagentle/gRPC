package main

import (
	"fmt"
	"log"

	pb "github.com/Kreagentle/gRPC/primenumbers/proto"
)

func (s *Server) CalculatePrimeNumbers(in *pb.PrimeRequest, stream pb.Calculator_CalculateServer) error {
	log.Printf("Function CalculatePrimeNumbers is invoked %v\n\n", in)
	for i := range []int{1, 2, 3, 4, 5} {
		data := fmt.Sprintf("Invokes %d the data %d", in.Number1, i)

		err := stream.Send(&pb.PrimeResponse{
			Number2: data,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
