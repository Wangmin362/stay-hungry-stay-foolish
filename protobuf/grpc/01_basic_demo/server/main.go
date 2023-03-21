package main

import (
	"context"
	"fmt"
	pb "github.com/golang/demo/protobuf/grpc/demo1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
	pb.UnimplementedReportServiceServer
}

func (s *HelloService) Hello(context.Context, *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Name: "hello server"}, nil
}

func main() {
	// 1. 监听端口
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen on 127.0.0.1:8888")

	// 2. 实例化gRPC服务端
	grpcServer := grpc.NewServer()

	// 3. 注册实现的服务实例
	service := &HelloService{}
	pb.RegisterReportServiceServer(grpcServer, service)

	// 4. 启动gRPC服务端
	fmt.Println("gRPC is running...")
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalf("gRPC server err:%s\n", err)
	}
}
