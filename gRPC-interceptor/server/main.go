package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	authpb "github.com/sevendollar/grpcintercept/proto/auth"
	hellopb "github.com/sevendollar/grpcintercept/proto/hello"
	worldpb "github.com/sevendollar/grpcintercept/proto/world"
)

func main() {
	store := newRedisStore("localhost:6379", "", "")
	defer store.close()
	jwtmanager := newJWTManager("secret", 0)
	authserver := newAuthServer(store, jwtmanager)
	interceptor := newAuthInterceptor(map[string][]string{}, jwtmanager)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.Unary(),
		),
	)

	authpb.RegisterAuthServiceServer(srv, authserver)
	hellopb.RegisterHelloServiceServer(srv, &HelloServer{})
	worldpb.RegisterWorldServiceServer(srv, &WorldServer{})
	reflection.Register(srv)

	go func() {
		log.Println("Server has started.")
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println()

	log.Println("Server is stopping...")
	srv.GracefulStop()
	log.Println("Server has stopped")
}

func m() {
	store := newRedisStore("localhost:6379", "", "")

	// u, err := newUser("joy@yahoo.com", "joyP@ssw0rd")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = store.save(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := store.addRole("joy@yahoo.com", "cisco")
	if err != nil {
		log.Fatal(err)
	}
}
