package main

import (
	"log"
	"net"

	"github.com/rubiagatra/grpc-go-node/go/pb/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":50000"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	messages.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
