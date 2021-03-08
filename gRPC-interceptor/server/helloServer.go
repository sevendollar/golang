package main

import (
	"context"

	hellopb "github.com/sevendollar/grpcintercept/proto/hello"
)

type HelloServer struct{}

func (hello *HelloServer) Man(context.Context, *hellopb.ManRequest) (*hellopb.ManResponse, error) {
	return &hellopb.ManResponse{
		Message: "hi",
	}, nil
}

func (hello *HelloServer) Cat(context.Context, *hellopb.CatRequest) (*hellopb.CatResponse, error) {
	return &hellopb.CatResponse{
		Message: "meow",
	}, nil
}

func (hello *HelloServer) Dog(context.Context, *hellopb.DogRequest) (*hellopb.DogResponse, error) {
	return &hellopb.DogResponse{
		Message: "bark",
	}, nil
}
