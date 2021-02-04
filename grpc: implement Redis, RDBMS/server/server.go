package main

import (
	"context"
	"fmt"
	"log"
	"time"

	hellopb "github.com/sevendollar/pb_test/proto"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {

	messageCH := make(chan string)
	var message string
	wait := make(chan struct{})

	go func() {
		defer close(messageCH)

		time.Sleep(time.Second * 5)

		messageCH <- "World"
	}()

	go func() {
		defer close(wait)

		message = <-messageCH
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
	}

	log.Println("[gRPC][HelloService][Hello][INFO] CALLER:", req.ProtoReflect().Descriptor().FullName())

	return &hellopb.HelloResponse{
		Message: message,
	}, nil

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

	// Send the values
	go getDBWithRedis(ctx, key, rCh, errCh)

	// Receive the values
	go func() {
		defer close(wait)
		err = <-errCh
		r = <-rCh

		// Notify when the task is done
		wait <- struct{}{}
	}()

	select {
	// context canceled
	case <-ctx.Done():
		log.Printf("[gRPC][HelloService][RedisGet][WARN] CALLER:%v REASON:%v\n", req.ProtoReflect().Descriptor().FullName(), ctx.Err())
		return nil, ctx.Err()

	// timeout
	case <-time.After(time.Second * 30):
		log.Println("[gRPC][HelloService][RedisGet][WARN] REASON:timeout")
		return nil, fmt.Errorf("timeout")

	// task done
	case <-wait:
	}

	if err != nil {
		log.Printf("[gRPC][HelloService][RedisGet][WARN] CALLER:%v REASON:%v\n", req.ProtoReflect().Descriptor().FullName(), err)
		return nil, err
	}
	log.Printf("[gRPC][HelloService][RedisGet][INFO] CALLER:%v, RESPONSE:%v\n", req.ProtoReflect().Descriptor().FullName(), r)

	return &hellopb.RedisGetResponse{
		Value: r,
	}, nil
}
