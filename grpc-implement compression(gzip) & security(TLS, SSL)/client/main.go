package main

import (
	"context"
	"fmt"
	"io"
	"log"

	hellopb "test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../cert/cert.pem", "")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(
		"localhost:50051",
		// implement security
		grpc.WithTransportCredentials(creds),
		// implement compression
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := hellopb.NewHelloServiceClient(conn)
	stream, err := c.HelloManyTimes(context.Background(), &hellopb.HelloManyTimesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	rlt := []byte{}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("bye")
				break
			}
			log.Fatal(err)
		}
		fmt.Println(resp.GetMessage())
		rlt = append(rlt, byte(resp.GetMessage()))
	}
	fmt.Println(string(rlt))
}
