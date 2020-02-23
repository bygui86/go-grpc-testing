package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/bygui86/go-grpc-testing/domain"
)

const (
	port = ":50051"
)

// GrpcServer is used to implement helloworld.GreeterServer.
type GrpcServer struct {
	proto.UnimplementedGreeterServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &GrpcServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *GrpcServer) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}
