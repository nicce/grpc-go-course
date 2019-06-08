package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	res := &calculatorpb.CalculatorResponse{
		Result: req.GetCalculator().GetValueOne() + req.GetCalculator().GetValueTwo(),
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberRequest, streamRes calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	number := req.GetNumber().GetNumber()
	var divider int32
	divider = 2

	for {
		if number > 1 {
			if number%divider == 0 {
				res := &calculatorpb.PrimeNumberResponse{
					Result: divider,
				}
				streamRes.Send(res)
				number = number / divider
			} else {
				divider = divider + 1
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

func (*server) ComputeAverage(streamReq calculatorpb.CalculatorService_ComputeAverageServer) error {
	count := 0
	sum := 0.0

	for {
		req, err := streamReq.Recv()

		if err == io.EOF {
			result := sum / float64(count)
			return streamReq.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error receving stream from client: %v", err)
		}

		sum = sum + req.GetNumber()
		count++
	}
}

func (*server) FindMaximum(streamReq calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("Find maximum invoked")

	max := int32(0)
	for {
		req, err := streamReq.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receving stream from client: %v", err)
			return err
		}

		if req.GetNumber() > max {
			max = req.GetNumber()
			err := streamReq.Send(&calculatorpb.FindMaximumResponse{
				Max: max,
			})
			if err != nil {
				log.Fatalf("Error sending stream: %v", err)
				return err
			}
		}

	}
}

func main() {
	fmt.Println("Starting calculator server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
