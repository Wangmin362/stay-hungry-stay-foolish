package main

import (
	"context"
	"fmt"
	pb "github.com/golang/demo/protobuf/grpc/demo1"
	"google.golang.org/grpc"
	"time"
)

func main() {
	// 1. 打开gRPC服务端链接
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	time.Sleep(20 * time.Second)

	// 2. 创建gRPC客户端
	client := pb.NewReportServiceClient(conn)

	// 4. 调用服务端提供的服务
	response, err := client.Hello(context.Background(), &pb.HelloReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Login Response: ", response)
}
