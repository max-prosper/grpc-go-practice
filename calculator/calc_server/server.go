package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/max-prosper/grpc-go-practice/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	intOne := req.GetIntOne()
	intTwo := req.GetIntTwo()
	result := intOne + intTwo

	res := &calcpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("I'm The Calculator")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterSumServceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
