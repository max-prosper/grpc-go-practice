package main

import (
	"fmt"
	"log"

	"github.com/max-prosper/grpc-go-practice/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.GreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
