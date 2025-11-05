package main

import (
	"log"
	"net"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"
	"google.golang.org/grpc"
)

const (
	Port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Starting gRPC server... at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
	log.Printf("gRPC server listening on %s", Port)
}
