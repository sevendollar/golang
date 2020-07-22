package main

import (
	"context"
	"fmt"
	"log"
	pb "proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	var (
		a int64 = 5
		b int64 = 12
	)

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	req := &pb.Request{
		A: a,
		B: b,
	}

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println()

	resp, err := client.Subtract(ctx, req)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
	fmt.Printf("a - b = %d\n", resp.GetRlt())

	resp, err = client.Add(ctx, req)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
	fmt.Printf("a + b = %d\n", resp.GetRlt())
}
