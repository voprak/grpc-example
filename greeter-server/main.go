// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	pb "github.com/voprak/grpc-example/greeter-server/pb"
	"github.com/voprak/grpc-example/greeter-server/rpci"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	log.Printf("Server running on port %v", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, rpci.NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
