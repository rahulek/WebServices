package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "math_service/proto"
)

var (
	port = flag.Int("port", 50051, "The Server Port")
)

type mathServer struct {
	pb.UnimplementedMathServiceServer
}

func (s *mathServer) Add(ctx context.Context, params *pb.TwoInts) (*pb.Result, error) {
	log.Printf("Add: N1=%v, N2=%v\n", params.N1, params.N2)
	return &pb.Result{
		Result: int64(params.N1 + params.N2),
	}, nil
}

func (s *mathServer) Multiply(ctx context.Context, params *pb.TwoInts) (*pb.Result, error) {
	log.Printf("Mulitply: N1= %v, N2= %v\n", params.N1, params.N2)
	return &pb.Result{
		Result: int64(params.N1 * params.N2),
	}, nil
}

func (s *mathServer) GetCapabilties(dummy *pb.Dummy, stream pb.MathService_GetCapabiltiesServer) error {
	log.Printf("GetCapabilities:")

	//Capabilities
	capabilties := [6]string{
		"Addition of Two Integers",
		"Addition of more than Two Integers",
		"Multiplication of Two Integers",
		"Multiplication of more than Two Integers",
		"Report Capabilities (This function)",
		"Transforming all the Integers passed-in",
	}

	for _, capability := range capabilties {
		if err := stream.Send(&pb.CapabilitiesResult{Capability: capability}); err != nil {
			return err
		}
	}

	return nil
}

func (s *mathServer) AddN(stream pb.MathService_AddNServer) error {
	sum := 0
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Result{
				Result: int64(sum),
			})
		}

		if err != nil {
			return err
		}

		log.Printf("AddN: Recevied %v\n", input.N1)
		sum += int(input.N1)
	}
}

func (s *mathServer) MultiplyN(stream pb.MathService_MultiplyNServer) error {
	mul := 1
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Result{
				Result: int64(mul),
			})
		}

		if err != nil {
			return err
		}

		log.Printf("MultiplyN: Recevied %v\n", input.N1)
		mul *= int(input.N1)
	}
}

func (s *mathServer) Transform(stream pb.MathService_TransformServer) error {
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		received := input.N1
		transformed := fmt.Sprintf("Hi, did you really send %v?", received)

		log.Printf("Transform: Received %v\n", received)
		log.Printf("Transform: Sending back %v\n", transformed)

		if err := stream.Send(&pb.TransformationResult{
			Transformed: transformed,
		}); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server listening on port: %v\n", *port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMathServiceServer(grpcServer, &mathServer{})
	grpcServer.Serve(lis)
}
