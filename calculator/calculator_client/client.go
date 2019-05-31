package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting calculator client..")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do a RPC..")

	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			ValueOne: 3,
			ValueTwo: 10,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed when invoking sum: %v", err)
	}

	log.Printf("Result from RPC call is: %v", res.GetResult())
}
