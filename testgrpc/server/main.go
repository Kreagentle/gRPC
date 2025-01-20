package main

import (
	"log"
	"net"
	
	pb "github.com/Kreagentle/gRPC/testgrpc/proto"
)

const address = "0.0.0.0:8085"

type Server struct {
	pb.Server
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen because of: %v\n", err)
	}

	log.Println("Listen on: %s\n", address)

	news := grpc.NewServer()
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to run gRPC server because of: %v\n", err)
	}
}
