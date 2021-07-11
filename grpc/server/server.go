package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dollarkillerx/light/codes"
	"github.com/dollarkillerx/light/cryptology"
	"github.com/dollarkillerx/light_rpc_benchmark/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var key = []byte("58a95a8f804b49e686f651a0d3f6e631")

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
		Rp:  coding([]byte("ok")),
		Msg: coding([]byte(fmt.Sprintf("resp: %s", decoding(message.Msg)))),
	}, nil
}

func coding(r []byte) []byte {
	encrypt, err := cryptology.AESEncrypt(key, r)
	if err != nil {
		log.Fatalln(err)
	}
	sn, _ := codes.CompressorManager.Get(codes.Snappy)
	zip, err := sn.Zip(encrypt)
	if err != nil {
		log.Fatalln(err)
	}
	return zip
}

func decoding(r []byte) []byte {
	sn, _ := codes.CompressorManager.Get(codes.Snappy)
	rc, err := sn.Unzip(r)
	if err != nil {
		log.Fatalln(err)
	}

	rb, err := cryptology.AESDecrypt(key, rc)
	if err != nil {
		log.Fatalln(err)
	}

	return rb
}
