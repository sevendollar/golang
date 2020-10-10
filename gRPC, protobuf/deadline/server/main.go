package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	calpb "github.com/sevendollar/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	// GLOBAL_RUN_DURATION ...
	// global timeout setting
	GLOBAL_RUN_DURATION = time.Duration(3000 * time.Millisecond)
)

type server struct{}

// implementatin of Hello()
func (*server) Hello(ctx context.Context, req *calpb.HelloRequest) (*calpb.HelloResponse, error) {
	log.Printf("Hello(), is being called... %v\n", req)

	return &calpb.HelloResponse{
		Message: "hello! " + req.GetName().GetFirstName() + " " + req.GetName().GetLastName(),
	}, nil
}

// implementatin of HelloWithTimeout()
func (*server) HelloWithTimeout(ctx context.Context, req *calpb.HelloWithTimeoutRequest) (*calpb.HelloWithTimeoutResponse, error) {
	log.Printf("HelloWithTimeout(), is being called... %v\n", req)

	run := make(chan struct{}, 1)

	// wait for certurn period of time to run the actual application
	go func(ctx context.Context) {
		// shows working status wrapped by goroutine
		go func(ctx context.Context) {
			defer fmt.Println("# goroutine exit properly.")
			for {
				select {
				case <-ctx.Done():
					if ctx.Err() == context.Canceled {
						return
					} else if ctx.Err() == context.DeadlineExceeded {
						return
					}
				default:
					fmt.Println("HelloWithTimeout(), working...")
				}
				time.Sleep(1 * time.Second)
			}
		}(ctx)

		// waiting period
		<-time.After(GLOBAL_RUN_DURATION)
		run <- struct{}{}
	}(ctx)

forLoop:
	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				log.Printf("HelloWithTimeout(), Client canceled: %v\n", ctx.Err())
				return nil, status.Errorf(codes.Canceled, "Client canceled: %v\n", ctx.Err())
			}
			if ctx.Err() == context.DeadlineExceeded {
				log.Printf("HelloWithTimeout(), Deadline exceeded: %v\n", ctx.Err())
				return nil, status.Errorf(codes.DeadlineExceeded, "Deadline exceeded: %v\n", ctx.Err())
			}
		case <-run:
			break forLoop
		}
	}

	result := &calpb.HelloWithTimeoutResponse{
		Message: "hello with timeout! " + req.GetName().GetFirstName() + " " + req.GetName().GetLastName(),
	}
	return result, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("can't listen to port: %v\n", err)
	}
	s := grpc.NewServer()
	calpb.RegisterCalServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		fmt.Print("gRPC service starting...\n\n")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error Servicing Server: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	fmt.Println("\ngracefully stopping server...")
	s.GracefulStop()
	fmt.Println("closing listener...")
	lis.Close()
	fmt.Println("gRPC service stopped!")
}
