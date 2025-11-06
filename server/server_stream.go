package main

import (
	"log"
	"time"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received request for server streaming: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name + " from gRPC Server Stream!",
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Simulate delay between sends
	}
	return nil
}
