package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc/chat"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	chat.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *chat.HelloRequest) (*chat.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &chat.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func main() {
	flag.Parsed()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	chat.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
