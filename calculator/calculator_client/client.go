package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"time"

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
	//doClientStreaming(c)
	doBiDiStreaming(c)
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

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	log.Printf("Starting to do a bi di streaming RPC...\n")

	reqList := []*calculatorpb.FindMaximumRequest{
		&calculatorpb.FindMaximumRequest{
			Number: int32(3),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(2),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(5),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(8),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(1),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(3),
		},
		&calculatorpb.FindMaximumRequest{
			Number: int32(9),
		},
	}

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		//Send requests

		for _, req := range reqList {
			log.Printf("Sending req: %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Errpr sending req: %v", err)
				close(waitc)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()

	go func() {
		// Receive data
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving stream: %v", err)
				break
			}

			log.Printf("Maximum is: %v \n", res.GetMax())

		}
		close(waitc)
	}()

	<-waitc
}
