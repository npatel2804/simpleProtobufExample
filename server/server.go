package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/npatel2804/simpleProtobufExample/server/protos"
	gRPC "google.golang.org/grpc"
)

type service struct{}

func (s *service) SumNum(ctx context.Context, r *pb.FunctionRequest) (*pb.FunctionResponse, error) {
	result := r.Num1 - r.Num2
	fmt.Println(result)
	//log.Fatalf(string(result))
	return &pb.FunctionResponse{Result: result}, nil
}

func (s *service) SubNum(ctx context.Context, r *pb.FunctionRequest) (*pb.FunctionResponse, error) {
	result := r.Num1 - r.Num2
	return &pb.FunctionResponse{Result: result}, nil
}

func main() {
	lst, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("starting Server....")
	grpc := gRPC.NewServer()
	fmt.Println("server is created")
	pb.RegisterCalculatorServer(grpc, &service{})
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		var ip net.IP
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
		fmt.Println(ip)
	}
	if err1 := grpc.Serve(lst); err1 != nil {
		panic(err)
	}

}
