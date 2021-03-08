package main

import (
	worldpb "github.com/sevendollar/grpcintercept/proto/world"
	"google.golang.org/grpc"
)

type WorldClient struct {
	service worldpb.WorldServiceClient
}

func newWorldClient(cc grpc.ClientConnInterface) *WorldClient {
	return &WorldClient{
		service: worldpb.NewWorldServiceClient(cc),
	}
}
