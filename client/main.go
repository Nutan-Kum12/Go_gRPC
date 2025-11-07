package main

import (
	"log"

	pb "github.com/Nutan-Kum12/Go_gRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"Akyu", "Bkyu", "Ckyu"},
	}
	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
