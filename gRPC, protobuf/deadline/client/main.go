package main

import (
	"context"
	"fmt"
	"log"
	"time"

	calpb "github.com/sevendollar/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to gRPC service: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		if conn.GetState() == connectivity.Shutdown {
			fmt.Println("\nconnection to gRPC closed!")
		} else {
			fmt.Println("\nERROR: connection to gRPC NOT closed!")
		}
	}(conn)
	defer conn.Close()

	c := calpb.NewCalServiceClient(conn)

	doHello(c)
	doHelloWithTimeout(c, 5*time.Second)

}

func doHelloWithTimeout(c calpb.CalServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.HelloWithTimeout(ctx, &calpb.HelloWithTimeoutRequest{
		Name: &calpb.People{
			FirstName: "Jeff",
			LastName:  "Lee",
		},
	})
	if err != nil {
		stautsErr, ok := status.FromError(err)
		if ok {
			// known grpc general codes`
			if stautsErr.Code() == codes.Canceled {
				log.Printf("Timeout! %v\n", stautsErr.Err())
			} else if stautsErr.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded! %v\n", stautsErr.Err())
			} else {
				log.Printf("Unknown gRPC code! %v\n", stautsErr.Err())
			}
			return
		}
		// unknown grpc codes
		log.Printf("Error! %v\n", err)
		return
	}

	fmt.Println(resp.GetMessage())
}

func doHello(c calpb.CalServiceClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := c.Hello(ctx, &calpb.HelloRequest{
		Name: &calpb.People{
			FirstName: "Jeff",
			LastName:  "Lee",
		},
	})
	if err != nil {
		log.Printf("Error! %v\n", err)
		return
	}

	fmt.Println(resp.GetMessage())
}
