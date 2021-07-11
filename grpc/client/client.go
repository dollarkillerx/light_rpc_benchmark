package main

import (
	"context"
	"log"

	"github.com/dollarkillerx/light_rpc_benchmark/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, e := grpc.Dial(":9001", grpc.WithInsecure()) // grpc.WithInsecure() 不安全的传输
	if e != nil {
		panic(e.Error())
	}
	client := proto.NewHelloClient(conn) // 注册上去

	_, err := client.Say(context.TODO(), &proto.BenchmarkMessage{
		Msg: "hello world",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
