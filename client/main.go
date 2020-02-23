package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	proto "github.com/bygui86/go-grpc-testing/domain"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	grpcConnection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcConnection.Close()
	grpcClient := proto.NewGreeterClient(grpcConnection)

	for {
		go greet(err, grpcClient)
		time.Sleep(3 * time.Second)
	}
}

func greet(err error, grpcClient proto.GreeterClient) {
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Contact the server and print out its response.
	r, err := grpcClient.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
