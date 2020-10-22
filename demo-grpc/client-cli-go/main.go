package main

import (
	greetpb "client/protobuf"
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	target := flag.String("host", "localhost", "host to listen to")
	targetPort := flag.String("port", "50051", "port to listen to")
	firstName := flag.String("firstname", "Golang", "first name")
	lastName := flag.String("lastname", "", "last name")
	flag.Parse()

	var otps []grpc.DialOption
	otps = append(otps, grpc.WithInsecure())

	conn, err := grpc.Dial((*target)+":"+(*targetPort), otps...)
	if err != nil {
		log.Fatalf("Cannot connecte to server, %v\n", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	resp, err := c.Hello(context.Background(), &greetpb.HelloRequest{
		PersonalInfo: &greetpb.PersonalInfo{
			FirstName: *firstName,
			LastName:  *lastName,
		},
	})
	if err != nil {
		log.Fatalf("Failed calling Helle() function, %v\n", err)
	}

	fmt.Println(resp.GetMessage())
}
