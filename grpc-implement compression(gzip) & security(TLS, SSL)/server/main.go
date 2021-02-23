package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	hellopb "test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	// implement compression
	_ "google.golang.org/grpc/encoding/gzip"
)

type server struct{}

func (*server) HelloManyTimes(req *hellopb.HelloManyTimesRequest, stream hellopb.HelloService_HelloManyTimesServer) error {
	s := "hello 世界!"

	for i := 0; i < len(s); i++ {
		err := stream.Send(&hellopb.HelloManyTimesResponse{
			Message: uint32(s[i]),
		})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(time.Millisecond * 500)
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// prepare the certificates
	creds, err := credentials.NewServerTLSFromFile("../cert/cert.pem", "../cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// implement security
	srv := grpc.NewServer(grpc.Creds(creds))

	hellopb.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	go func() {
		fmt.Println("server running...")
		if err := srv.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("server stopping...")
	srv.GracefulStop()
	fmt.Println("server stopped")

}
