package main

import (
	"context"
	"fmt"
	"net"

	"github.com/dollarkillerx/light_rpc_benchmark/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, e := net.Listen("tcp", ":9001")
	if e != nil {
		panic(e.Error())
	}

	// 创建一个没有注册的grpc
	srv := grpc.NewServer()
	proto.RegisterHelloServer(srv, &server{})
	reflection.Register(srv)

	// 监听
	if e := srv.Serve(listener); e != nil {
		panic(e.Error())
	}
}

type server struct{}

func (s server) Say(ctx context.Context, message *proto.BenchmarkMessage) (*proto.BenchmarkMessage, error) {
	return &proto.BenchmarkMessage{
		Msg: fmt.Sprintf("resp: %s", message.Msg),
	}, nil
}
