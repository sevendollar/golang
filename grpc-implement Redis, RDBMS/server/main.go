package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	hellopb "github.com/sevendollar/pb_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	conn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal((err))
	}
	defer conn.Close()

	srv := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(srv, &server{})
	reflection.Register(srv)

	go func() {
		fmt.Println("gRPC Servcie Is Started!")
		if err := srv.Serve(conn); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("\ngRPC Servcie Is Stopping...\n")

	srv.GracefulStop()
	fmt.Println("gRPC Servcie Has Gracefully Stopped!")
}
