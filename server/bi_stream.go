package main

import (
	"io"
	"log"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name) //server side log
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name, //client side response
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
