package main

import (
	"github.com/nikit34/basic_tutorial/grpc/simple/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &proto.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}