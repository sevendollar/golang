package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	hellopb "temp2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {

	r := &hellopb.HelloResponse{
		Message: "Hello World",
	}

	wait := make(chan struct{})

	// the task
	go func() {
		time.Sleep(time.Second * 10)
		wait <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Printf("[gRPC][HelloService_Hello][WARN] %v\n", ctx.Err())
		return &hellopb.HelloResponse{}, ctx.Err()

	case <-time.After(time.Second * 6):
		log.Printf("[gRPC][HelloService_Hello][WARN] %v\n", fmt.Errorf("Time Out"))
		return &hellopb.HelloResponse{}, fmt.Errorf("Time Out")

	case <-wait:
		// task done!
	}

	return r, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("gRPC Service Started!")
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Gracefully Shutdown The gRPC Service
	srv.GracefulStop()
	close(quit)
	fmt.Println()
	fmt.Println("gRPC Service Gracefully Shutdown!")

}
