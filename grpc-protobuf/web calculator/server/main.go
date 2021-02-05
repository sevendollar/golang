package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "proto/webCalculator"
)

func main() {
	listenPort := 2343

	listener, err := net.Listen("tcp", ":"+fmt.Sprintf("%d", listenPort))
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterCalculatorServer(srv, &calculator{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

}
