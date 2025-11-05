package main

import (
	"log"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"

	"context"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from SayHello: %v", res.Message)
}
