package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	hellopb "github.com/sevendollar/grpcintercept/proto/hello"
	worldpb "github.com/sevendollar/grpcintercept/proto/world"
)

func main() {
	cc1, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cc1.Close()

	authclient := newAuthClient(cc1, "joy@yahoo.com", "joyP@ssw0rd")
	authinterceptor := newAuthInterceptor(authclient)

	cc2, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			authinterceptor.Unary(),
			helloInterceptorUnary(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cc2.Close()

	ctx := context.Background()

	helloservice := newHelloClient(cc2)
	// Man
	manRlt, err := helloservice.service.Man(ctx, &hellopb.ManRequest{})
	if err != nil {
		log.Println("[Man][ERROR]:", err)
	} else {
		fmt.Println("Man:", manRlt.GetMessage())
	}

	// Cat
	catRlt, err := helloservice.service.Cat(ctx, &hellopb.CatRequest{})
	if err != nil {
		log.Println("[Cat][ERROR]:", err)
	} else {
		fmt.Println("Cat:", catRlt.GetMessage())
	}

	// Dog
	dogRlt, err := helloservice.service.Dog(ctx, &hellopb.DogRequest{})
	if err != nil {
		log.Println("[Dog][ERROR]:", err)
	} else {
		fmt.Println("Dog:", dogRlt.GetMessage())
	}

	worldservice := newWorldClient(cc2)
	// Tuna
	tunaRlt, err := worldservice.service.Tuna(ctx, &worldpb.TunaRequest{})
	if err != nil {
		log.Println("[Tuna][ERROR]:", err)
	} else {
		fmt.Println("Tuna:", tunaRlt.GetMessage())
	}

	// Beef
	beefRlt, err := worldservice.service.Beef(ctx, &worldpb.BeefRequest{})
	if err != nil {
		log.Println("[Beef][ERROR]:", err)
	} else {
		fmt.Println("Beef:", beefRlt.GetMessage())
	}

}
