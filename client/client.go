package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/npatel2804/simpleProtobufExample/client/protos"
	"google.golang.org/grpc"
)

func main() {
	//stores space separated values into successive arguments
	time.Sleep(40 * time.Second)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	fmt.Println(err)
	c := pb.NewCalculatorClient(conn)
	rq := pb.FunctionRequest{Num1: 9, Num2: 4}
	var storageVariable string
	fmt.Println("Choose 1 or 2")
	fmt.Scan(&storageVariable) //assuming fmt is imported

	if storageVariable == "1" {
		//fmt.Printf("1")
		ctx, err1 := context.WithTimeout(context.Background(), 20*time.Second)
		fmt.Printf("%v", err1)
		rp, err2 := c.SumNum(ctx, &rq)
		fmt.Printf("%v", err2)
		fmt.Print(rp.GetResult())
	} else {
		ctx, err1 := context.WithTimeout(context.Background(), 20*time.Second)
		fmt.Print(err1)
		rp, err2 := c.SubNum(ctx, &rq)
		fmt.Print(err2)
		fmt.Print(rp.GetResult())
		fmt.Printf("2")
	}
}
