package main

import (
	"log"

	pb "github.com/Kreagentle/gRPC/average/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "0.0.0.0:8085"

func main() {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		if err != nil {
			log.Printf("Failed to close connection: %v\n", err)
		}
	}(connection)

	log.Printf("Connected to: %s\n", address)

	client := pb.NewCalculatorClient(connection)
	// AverageCalculate(client)
}
