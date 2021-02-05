package main

import (
	"context"
	pb "proto/webCalculator"
)

type calculator struct{}

func (c *calculator) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	rlt := a + b
	return &pb.Response{Result: rlt}, nil
}

func (c *calculator) Subtract(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	rlt := a - b
	return &pb.Response{Result: rlt}, nil
}

func (c *calculator) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	rlt := a * b
	return &pb.Response{Result: rlt}, nil
}

func (c *calculator) Divid(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	rlt := a / b
	return &pb.Response{Result: rlt}, nil
}
