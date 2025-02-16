package main

import (
	"log"

	pb "github.com/Kreagentle/gRPC/primenumbers/proto"
)

func (s *Server) CalculatePrimeNumbers(in *pb.PrimeRequest, stream pb.Calculator_CalculateServer) error {
	log.Printf("Function CalculatePrimeNumbers is invoked %v\n\n", in)

	n := in.Number1
	div := int64(2)

	for n > 1 {
		if n%div == 0 {
			err := stream.Send(&pb.PrimeResponse{
				Number2: div,
			})
			if err != nil {
				return err
			}
			n /= div
		} else {
			div += 1
		}
	}

	return nil
}
