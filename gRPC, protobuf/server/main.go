package main

import (
	"context"
	"log"
	"net"
	pb "proto"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type calculator struct{}

func (c *calculator) Add(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	a, b := r.GetA(), r.GetB()
	rlt := a + b
	return &pb.Response{Rlt: rlt}, nil
}

func (c *calculator) Subtract(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	a, b := r.GetA(), r.GetB()
	rlt := a - b
	return &pb.Response{Rlt: rlt}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	srv := grpc.NewServer()

	pb.RegisterCalculatorServer(srv, &calculator{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
}
