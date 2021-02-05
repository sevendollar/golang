package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	dummypb "github.com/sevendollar/test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Can't listen to the service: %v\n", err)
	}
	srv := grpc.NewServer()
	dummypb.RegisterDummyServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("start gRPC service...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Oops, something went wrong when servicing the service, %v\n", err)
	}

}

type server struct{}

func (*server) Hello(context.Context, *dummypb.HelloRequest) (*dummypb.HelloResponse, error) {
	return &dummypb.HelloResponse{
		Result: "hello gRPC!",
	}, nil
}

func (*server) BioHello(stream dummypb.DummyService_BioHelloServer) error {
	wg := sync.WaitGroup{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can't recieve the client streaming: %v\n", err)
			break
		}

		wg.Add(1)

		go func() {
			defer wg.Done()

			rand.Seed(time.Now().UnixNano())
			n := time.Duration(rand.Int63n(10000))
			time.Sleep(n * time.Millisecond)

			if err := stream.Send(&dummypb.BioHelloResponse{
				Message: "SERVER: hello " + req.GetName() + "!",
			}); err != nil {
				log.Fatalf("problem sending the server streaming: %v\n", err)
			}
		}()

	}
	wg.Wait()

	return nil
}

func (*server) HelloManyTimes(req *dummypb.HelloManyTimesRequest, stream dummypb.DummyService_HelloManyTimesServer) error {
	for i := 0; i <= 10; i++ {
		err := stream.Send(&dummypb.HelloManyTimesResponse{
			Result: "hello" + strconv.Itoa(i),
		})
		if err != nil {
			log.Println("Can't stream data to client, ", err)
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongHello(stream dummypb.DummyService_LongHelloServer) error {
	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := stream.SendAndClose(&dummypb.LongHelloResponse{
				Result: "Hello " + result,
			}); err != nil {
				log.Fatalf("somthing went wrong when sending response to the client, %v\n", err)
			}
			return nil
		}
		if err != nil {
			log.Fatalf("can't read the stream coming from the client, %v\n", err)
			return err
		}

		name := req.GetFirstName() + " " + req.GetLastName() + "! "
		result += name
	}

}
