package main

import (
	"log"
	"net"

	pb "github.com/Kreagentle/gRPC/calculator/proto"

	"google.golang.org/grpc"
)

const address = "0.0.0.0:8085"

type Server struct {
	pb.CalculatorServer
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen because of: %v\n", err)
	}

	log.Printf("Listen on: %s\n", address)

	news := grpc.NewServer()
	pb.RegisterCalculatorServer(news, &Server{})
	err = news.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to run gRPC server because of: %v\n", err)
	}
}
