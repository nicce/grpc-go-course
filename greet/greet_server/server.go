package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function invoked with: %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	res := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
