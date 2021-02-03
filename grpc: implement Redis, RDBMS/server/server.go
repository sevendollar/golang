package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	hellopb "github.com/sevendollar/pb_test/proto"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	r := &hellopb.HelloResponse{
		Message: "world",
	}

	wait := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)

		wait <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Printf("[gRPC][HelloService][Hello][WARN] CALLER:%v REASON:%v\n", req.ProtoReflect().Descriptor().FullName(), ctx.Err())
		return nil, ctx.Err()

	case <-time.After(time.Second * 30):
		log.Println("[gRPC][HelloService][Hello][WARN] REASON:timeout")
		return nil, fmt.Errorf("timeout")

	case <-wait:
		log.Println("[gRPC][HelloService][Hello][INFO] CALLER:", req.ProtoReflect().Descriptor().FullName())
		close(wait)

	}

	return r, nil

}

func (*server) Ping(ctx context.Context, req *hellopb.PingRequest) (*hellopb.PingResponse, error) {
	log.Println("[gRPC][HelloService][Ping][INFO] CALLER:", req.ProtoReflect().Descriptor().FullName())
	return &hellopb.PingResponse{
		Message: "pong!",
	}, nil
}

func (*server) RedisGet(ctx context.Context, req *hellopb.RedisGetRequest) (*hellopb.RedisGetResponse, error) {
	key := req.GetKey()
	if key == "" {
		log.Println("[gRPC][HelloService][RedisGet][INFO] CALLER:", req.ProtoReflect().Descriptor().FullName())
		return nil, fmt.Errorf("missing key")
	}

	wait := make(chan struct{})
	rCh, errCh := make(chan string), make(chan error)
	var r string
	var err error

	go func() {
		c := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		defer c.Close()

		v, err := c.Get(ctx, key).Result()
		errCh <- err
		close(errCh)
		rCh <- v
		close(rCh)
	}()

	go func() {
		err = <-errCh
		r = <-rCh

		wait <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Printf("[gRPC][HelloService][RedisGet][WARN] CALLER:%v REASON:%v\n", req.ProtoReflect().Descriptor().FullName(), ctx.Err())
		return nil, ctx.Err()

	case <-time.After(time.Second * 30):
		log.Println("[gRPC][HelloService][RedisGet][WARN] REASON:timeout")
		return nil, fmt.Errorf("timeout")

	case <-wait:
		close(wait)
	}

	if err != nil {
		log.Printf("[gRPC][HelloService][RedisGet][WARN] CALLER:%v REASON:%v\n", req.ProtoReflect().Descriptor().FullName(), err)
		return nil, err
	}
	log.Printf("[gRPC][HelloService][RedisGet][INFO] CALLER:%v\n", req.ProtoReflect().Descriptor().FullName())

	return &hellopb.RedisGetResponse{
		Value: r,
	}, nil
}
