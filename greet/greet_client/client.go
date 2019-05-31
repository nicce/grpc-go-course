package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client..")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	// doUnary(c)
	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Niclas",
			LastName:  "Fredriksson",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greet: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Niclas",
			LastName:  "Fredriksson",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to invoke GreetManyTimes %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// We've reached end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}
