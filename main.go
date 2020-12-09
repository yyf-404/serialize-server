package main

import (
	"fmt"
	serialize "github.com/yyf-404/serialize_pb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"serialize-server/service"
)

func main() {
	// 监听本地的8972端口

	lis, err := net.Listen("tcp", ":8972")
	fmt.Println("init service")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	serialize.RegisterSerializeServiceServer(s, &service.SerializeService{})
	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
