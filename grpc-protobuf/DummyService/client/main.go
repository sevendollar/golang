package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	dummypb "github.com/sevendollar/client/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't dial to gRPC server, %v\n", err)
	}
	defer conn.Close()

	c := dummypb.NewDummyServiceClient(conn)
	// hello(c)
	// getStream(c)
	// streamToServer(c)
	bioStream(c)
}

func bioStream(c dummypb.DummyServiceClient) {
	stream, err := c.BioHello(context.Background())
	if err != nil {
		log.Fatalf("Error when connecting to server: %v\n", err)
		return
	}

	done := make(chan struct{})
	nameList := []string{
		"1 tune",
		"2 fish",
		"3 swordfish",
		"4 starfish",
		"5 jellyfish",
		"6 seaturtle",
		"7 beef",
		"8 pork",
		"9 chiken",
	}
	names := []*dummypb.BioHelloRequest{}

	for _, v := range nameList {
		x := &dummypb.BioHelloRequest{
			Name: v,
		}
		names = append(names, x)
	}

	// names := []*dummypb.BioHelloRequest{
	// 	&dummypb.BioHelloRequest{
	// 		Name: "jef",
	// 	},
	// 	&dummypb.BioHelloRequest{
	// 		Name: "peter",
	// 	},
	// 	&dummypb.BioHelloRequest{
	// 		Name: "sam",
	// 	},
	// 	&dummypb.BioHelloRequest{
	// 		Name: "john",
	// 	},
	// 	&dummypb.BioHelloRequest{
	// 		Name: "dounal",
	// 	},
	// }

	// send
	go func() {
		for _, n := range names {
			fmt.Printf("CLIENT: sending message to server: %v\n", n)

			if err := stream.Send(n); err != nil {
				log.Fatalf("Error sending streaming to server: %v\n", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// recieve
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error recieving streaming from server: %v\n", err)
				break
			}
			fmt.Println(resp.GetMessage())
		}
		close(done)
	}()

	<-done

}

func streamToServer(c dummypb.DummyServiceClient) {

	names := []*dummypb.LongHelloRequest{
		&dummypb.LongHelloRequest{
			FirstName: "jef",
			LastName:  "lee",
		},
		&dummypb.LongHelloRequest{
			FirstName: "peter",
			LastName:  "hush",
		},
		&dummypb.LongHelloRequest{
			FirstName: "sam",
			LastName:  "smith",
		},
		&dummypb.LongHelloRequest{
			FirstName: "john",
			LastName:  "locker",
		},
		&dummypb.LongHelloRequest{
			FirstName: "dounal",
			LastName:  "trump",
		},
	}
	stream, err := c.LongHello(context.Background())
	if err != nil {
		log.Fatalf("can't stream to server, %v\n", err)
	}

	for _, v := range names {
		fmt.Printf("streaming <%v> to server...\n", v)
		err := stream.Send(v)
		if err != nil {
			log.Fatalf("can't stream to server, %v\n", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	req, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("can't stream to server, %v\n", err)
	}
	fmt.Println(req.GetResult())

}

func getStream(c dummypb.DummyServiceClient) {
	stream, err := c.HelloManyTimes(context.Background(), &dummypb.HelloManyTimesRequest{})
	if err != nil {
		log.Fatalf("something went wrong with SERVER STREAMING, %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something went wrong when reading stream, %v\n", err)
		}
		fmt.Println(msg.GetResult())
	}
}

func hello(c dummypb.DummyServiceClient) {
	result, err := c.Hello(context.Background(), &dummypb.HelloRequest{})
	if err != nil {
		log.Fatalf("calling Hello() went wrong, %v\n", err)
	}
	fmt.Println(result.GetResult())
}
