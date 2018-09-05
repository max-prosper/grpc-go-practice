package main

import (
	"context"
	"fmt"
	"log"

	"github.com/max-prosper/grpc-go-practice/calculator/calcpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a Sum service client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewSumServceClient(cc)
	doSum(c)
}

func doSum(c calcpb.SumServceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calcpb.SumRequest{
		IntOne: 5,
		IntTwo: 8,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response form Sum: %v", res.Result)
}
