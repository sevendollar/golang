package main

import (
	"context"
	"log"
	"net"

	pb "server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type merge struct{}

func (m *merge) Do(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	rlt := new(pb.Response)
	a := r.GetA()
	b := r.GetB()

	rlt.Message = b + " " + a

	return rlt, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	srv := grpc.NewServer()

	pb.RegisterMergeServer(srv, &merge{})
	reflection.Register(srv)

	log.Println("gRPC starting at TCP:8080")
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
