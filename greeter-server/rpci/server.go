// Package rpci implements a server for Greeter service.
package rpci

import (
	"context"
	"log"

	pb "github.com/voprak/grpc-example/greeter-server/pb"
)

const (
	port = ":9000"
)

// server is used to implement greeter.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func NewServer() pb.GreeterServer {
	return &server{}
}

// SayHello implements greeter.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
