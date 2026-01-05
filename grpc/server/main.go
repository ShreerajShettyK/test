package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-example/proto"

	"google.golang.org/grpc"
)

// Server struct implements the Calculator service
type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

// Add implements the Add RPC method
func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received Add request: a=%d, b=%d", req.A, req.B)

	result := req.A + req.B

	return &pb.AddResponse{
		Result: result,
	}, nil
}

// Multiply implements the Multiply RPC method
func (s *calculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("Received Multiply request: a=%d, b=%d", req.A, req.B)

	result := req.A * req.B

	return &pb.MultiplyResponse{
		Result: result,
	}, nil
}

// Fibonacci implements the server-streaming Fibonacci RPC method
func (s *calculatorServer) Fibonacci(req *pb.FibonacciRequest, stream pb.Calculator_FibonacciServer) error {
	log.Printf("Received Fibonacci request: n=%d", req.N)

	a, b := 0, 1

	for i := int32(0); i < req.N; i++ {
		// Send each fibonacci number to the client
		if err := stream.Send(&pb.FibonacciResponse{Value: int32(a)}); err != nil {
			return err
		}

		// Calculate next fibonacci number
		a, b = b, a+b

		// Add small delay to demonstrate streaming
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func main() {
	// Create TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register our calculator service with the gRPC server
	pb.RegisterCalculatorServer(grpcServer, &calculatorServer{})

	fmt.Println("🚀 gRPC Server started on port 50051")
	fmt.Println("Waiting for client connections...")

	// Start serving requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
