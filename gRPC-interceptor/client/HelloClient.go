package main

import (
	"google.golang.org/grpc"

	hellopb "github.com/sevendollar/grpcintercept/proto/hello"
)

type HelloClient struct {
	service hellopb.HelloServiceClient
}

func newHelloClient(cc grpc.ClientConnInterface) *HelloClient {
	return &HelloClient{
		service: hellopb.NewHelloServiceClient(cc),
	}
}
