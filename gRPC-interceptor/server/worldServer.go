package main

import (
	"context"

	worldpb "github.com/sevendollar/grpcintercept/proto/world"
)

type WorldServer struct{}

func (world *WorldServer) Tuna(context.Context, *worldpb.TunaRequest) (*worldpb.TunaResponse, error) {
	return &worldpb.TunaResponse{
		Message: "fish",
	}, nil
}
func (world *WorldServer) Beef(context.Context, *worldpb.BeefRequest) (*worldpb.BeefResponse, error) {
	return &worldpb.BeefResponse{
		Message: "roll",
	}, nil
}
