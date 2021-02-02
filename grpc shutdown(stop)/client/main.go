package main

import (
	"context"
	"fmt"
	"log"

	hellopb "tmp3/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := hellopb.NewHelloServiceClient(conn)
	r, err := c.Hello(context.Background(), &hellopb.HelloRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.GetMessage())

}
