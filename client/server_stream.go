package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Error calling SayHelloServerStream: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving from SayHelloServerStream: %v", err)
		}
		log.Printf("Received response from SayHelloServerStream: %v", resp)
	}
}
