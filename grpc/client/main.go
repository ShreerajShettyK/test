package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "grpc-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a Calculator client
	client := pb.NewCalculatorClient(conn)

	fmt.Println("✅ Connected to gRPC server")
	fmt.Println()

	// Example 1: Unary RPC - Add
	fmt.Println("=== Example 1: Add (Unary RPC) ===")
	addRequest := &pb.AddRequest{A: 10, B: 20}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addResponse, err := client.Add(ctx, addRequest)
	if err != nil {
		log.Fatalf("Add failed: %v", err)
	}
	fmt.Printf("Add(%d, %d) = %d\n", addRequest.A, addRequest.B, addResponse.Result)
	fmt.Println()

	// Example 2: Unary RPC - Multiply
	fmt.Println("=== Example 2: Multiply (Unary RPC) ===")
	multiplyRequest := &pb.MultiplyRequest{A: 5, B: 7}
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	multiplyResponse, err := client.Multiply(ctx2, multiplyRequest)
	if err != nil {
		log.Fatalf("Multiply failed: %v", err)
	}
	fmt.Printf("Multiply(%d, %d) = %d\n", multiplyRequest.A, multiplyRequest.B, multiplyResponse.Result)
	fmt.Println()

	// Example 3: Server Streaming RPC - Fibonacci
	fmt.Println("=== Example 3: Fibonacci (Server Streaming RPC) ===")
	fibRequest := &pb.FibonacciRequest{N: 10}
	ctx3, cancel3 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel3()

	stream, err := client.Fibonacci(ctx3, fibRequest)
	if err != nil {
		log.Fatalf("Fibonacci failed: %v", err)
	}

	fmt.Printf("First %d Fibonacci numbers:\n", fibRequest.N)
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// Stream finished
			break
		}
		if err != nil {
			log.Fatalf("Error receiving from stream: %v", err)
		}
		fmt.Printf("%d ", response.Value)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("✅ All examples completed successfully!")
}
