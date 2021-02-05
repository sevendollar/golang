package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "client/proto"

	"google.golang.org/grpc"
)

func main() {
	var (
		a        = "hello"
		b        = "Golang"
		dialHost = ""
		dialPort = ""
	)

	flag.StringVar(&dialHost, "host", "localhost", "dial information")
	flag.StringVar(&dialPort, "port", "8080", "dial information")
	flag.Parse()

	if tmpHost := os.Getenv("host"); tmpHost != "" {
		dialHost = tmpHost
	}
	if tmpPort := os.Getenv("port"); tmpPort != "" {
		dialPort = tmpPort
	}
	// fmt.Println(dialHost)
	// fmt.Println(dialPort)

	dialInfo := dialHost + ":" + dialPort
	// fmt.Println(dialInfo)

	conn, err := grpc.Dial(dialInfo, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewMergeClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rlt, err := client.Do(ctx, &pb.Request{A: a, B: b})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(rlt.GetMessage())

}
