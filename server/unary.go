package main

import (
	"context"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from gRPC Server!"}, nil
}
