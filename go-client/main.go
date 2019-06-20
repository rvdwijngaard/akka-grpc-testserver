package main

import (
	"context"
	"fmt"
	pb "grpc-tester/proto"
	"log"
	"time"

	"github.com/pkg/errors" //f
	"google.golang.org/grpc"
)

func main() {
	client, err := NewRuntimeClient("localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Mark!"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response %+v", resp)
}

func NewRuntimeClient(target string) (pb.GreeterServiceClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "could not get monitoring grpc client at %s", target)
	}
	return pb.NewGreeterServiceClient(conn), nil
}
