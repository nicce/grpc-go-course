package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
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

	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do a server streaming RPC...")
	req := &calculatorpb.PrimeNumberRequest{
		Number: &calculatorpb.PrimeNumber{
			Number: 120,
		},
	}

	streamRes, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to do PrimeNumberDecomp:  %v", err)
	}

	for {
		msg, err := streamRes.Recv()
		if err == io.EOF {
			// We've received the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error streaming data: %v", err)
		}

		log.Printf("Result from stream: %v", msg)
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do a client streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())

	for i := 1; i <= 400; i++ {
		req := &calculatorpb.ComputeAverageRequest{
			Number: float64(i),
		}

		if err != nil {
			log.Fatalf("Error while calling ComputeAverage: %v", err)
		}

		log.Printf("Sending: %v", req)

		stream.Send(req)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receving response: %v", err)
	}

	log.Printf("Compute average response: %v", res)
}
