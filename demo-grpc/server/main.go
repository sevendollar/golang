package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/reflection"

	greetpb "github.com/sevendollar/protobuf/grpc-server/protobuf"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, req *greetpb.HelloRequest) (*greetpb.HelloResponse, error) {
	firstname := req.GetPersonalInfo().GetFirstName()
	lastname := req.GetPersonalInfo().GetLastName()

	res := "Hello, " + firstname + " " + lastname

	rlt := &greetpb.HelloResponse{
		Message: res,
	}

	log.Printf("Client calling, %v\n", req)
	return rlt, nil
}

func main() {
	port := flag.String("port", "50051", "port to listen to.")
	flag.Parse()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Cannot listen to tcp:%s %v\n", *port, err)
	}

	srv := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(srv, &server{})
	reflection.Register(srv)

	go func() {
		fmt.Println("gRPC server listening to :" + *port)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Cannot serve the gRPC server %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c

	fmt.Println("\nserver stopped")
	srv.GracefulStop()

	fmt.Println("stop listening.")
	lis.Close()
}
