package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
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

	doUnary(c)

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
