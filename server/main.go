package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	Port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
	log.Printf("gRPC server listening on %s", Port)
}
