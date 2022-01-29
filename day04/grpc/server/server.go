package main

import (
	"context"
	"fmt"
	pb "grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Println("sending result:", result)

	res := &pb.AddResponse{
		Result: result,
	}
	return res, nil
}
func (s *server) GeneratePrimes(req *pb.PrimeRequest, stream pb.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Request receive for generating prime from %d to %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			res := &pb.PrimeResponse{
				PrimeNumber: no,
			}
			fmt.Println("Sending prime no : ", no)
			stream.Send(res)
		}
	}
	return nil
}
func (s *server) CalculateAverage(stream pb.AppService_CalculateAverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Received %d for calculating average\n", req.GetNo())
		sum += req.GetNo()
		count++
	}
	avg := sum / count
	res := &pb.AverageResponse{
		Average: avg,
	}
	fmt.Println("Sending average : ", avg)
	stream.SendAndClose(res)
	return nil
}

func (s *server) GenerateMultiPrimes(stream pb.AppService_GenerateMultiPrimesServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		start := req.GetStart()
		end := req.GetEnd()
		for no := start; no < end; no++ {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				res := &pb.PrimeResponse{
					PrimeNumber: no,
				}
				fmt.Println("Sending prime no : ", no)
				stream.Send(res)
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	var s = &server{}
	grpcServer := grpc.NewServer()
	pb.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}

func isPrime(no int32) bool {
	if no <= 0 {
		return false
	}
	if no <= 3 {
		return true
	}
	for i := int32(2); i <= no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
