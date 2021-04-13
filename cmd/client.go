package main

import (
	"context"
	"fmt"
	"gokit/test/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewTestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/****** обычные message message***/
	/* 	r, err := c.SayHello(ctx, &dt.HelloRequest{Name: "Hi from client"})
	   	p, err := c.GetById(ctx, &dt.Product_Id{Value: 1})
	*/
	r, err := c.Validate(ctx, &pb.Input{Str: "[()]{}{[()()]()}"})
	f, err := c.Fix(ctx, &pb.Input{Str: "[()]{}{[()()]()}[[["})

	fmt.Println(r)
	fmt.Println(f)
}
