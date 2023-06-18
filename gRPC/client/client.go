package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "math_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewMathServiceClient(conn)

	addTwo(client, 10, 20)
	addTwo(client, -100, 145)

	mulTwo(client, 1000, -6)
	mulTwo(client, 25, 50)

	streamedAdd(client, 10, 20, 30, 40, 50)
	streamedAdd(client, -10, -20, 300, 400, 500)

	streamedMultiply(client, 30, 40, -10)
	streamedMultiply(client, 12, 40, 2)

	capabilities(client)

	transformation(client, 11, 111, 1111)
	transformation(client, -9, -10, -11)
}

// Add Two Ints
func addTwo(client pb.MathServiceClient, n1 int32, n2 int32) error {
	res, err := client.Add(context.Background(), &pb.TwoInts{N1: n1, N2: n2})
	if err != nil {
		return err
	}
	fmt.Printf("addTwo: n1= %v, n2= %v, Result: %v\n", n1, n2, res.Result)
	return nil
}

// Multiply Two Ints
func mulTwo(client pb.MathServiceClient, n1 int32, n2 int32) error {
	res, err := client.Multiply(context.Background(), &pb.TwoInts{N1: n1, N2: n2})
	if err != nil {
		return err
	}
	fmt.Printf("mulTwo: n1= %v, n2= %v, Result: %v\n", n1, n2, res.Result)
	return nil
}

// Streaming Addition
func streamedAdd(client pb.MathServiceClient, ints ...int) error {
	stream, err := client.AddN(context.Background())
	if err != nil {
		return err
	}

	for _, n := range ints {
		if err := stream.Send(&pb.IntInput{N1: int32(n)}); err != nil {
			log.Printf("streamedAdd: Failed to send data: %v\n", err)
			return err
		}
		fmt.Printf("(Sent %v) ", n)
	}
	fmt.Println()

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("streamedAdd: Could not get result: %v\n", err)
	}
	fmt.Printf("streamedAdd: Result: %v\n", reply.Result)

	return nil
}

// Streaming multiplication
func streamedMultiply(client pb.MathServiceClient, ints ...int) error {
	stream, err := client.MultiplyN(context.Background())
	if err != nil {
		return err
	}

	for _, n := range ints {
		if err := stream.Send(&pb.IntInput{N1: int32(n)}); err != nil {
			log.Printf("streamedMultiply: Failed to send data: %v\n", err)
			return err
		}
		fmt.Printf("(Sent %v) ", n)
	}
	fmt.Println()

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("streamedultiply: Could not get result: %v\n", err)
	}
	fmt.Printf("streamedMultiply: Result: %v\n", reply.Result)

	return nil
}

// Get Server Capabilities
func capabilities(client pb.MathServiceClient) error {
	stream, err := client.GetCapabilties(context.Background(), &pb.Dummy{})
	if err != nil {
		fmt.Printf("Capabilities: Failed to get the capabilities: %v\n", err)
		return err
	}

	for {
		cap, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Capabilities: Error getting data from the stream: %v\n", err)
			return err
		}
		fmt.Printf("Capabilities: Received: %v\n", cap.Capability)
	}

	return nil
}

// Two Way Transformation
func transformation(client pb.MathServiceClient, ints ...int) error {
	stream, err := client.Transform(context.Background())
	if err != nil {
		fmt.Printf("transformation: Failed to get the stream: %v\n", err)
		return err
	}

	waitCh := make(chan struct{})
	go func() {
		for {
			transformed, err := stream.Recv()
			if err == io.EOF {
				close(waitCh)
				return
			}
			if err != nil {
				fmt.Printf("transformation: Error: %v\n", err)
				return
			}
			fmt.Printf("transformation: Received: %v\n", transformed.Transformed)
		}
	}()

	for _, n := range ints {
		if err := stream.Send(&pb.IntInput{N1: int32(n)}); err != nil {
			fmt.Printf("transformation: Failed to send %v\n", n)
		}
		fmt.Printf("transformation: [Sent %v]\n", n)
	}
	stream.CloseSend()

	<-waitCh

	return nil
}
