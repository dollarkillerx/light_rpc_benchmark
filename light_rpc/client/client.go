package main

import (
	"fmt"
	"log"

	"github.com/dollarkillerx/light"
	"github.com/dollarkillerx/light/client"
	"github.com/dollarkillerx/light/discovery"
	"github.com/dollarkillerx/light/transport"
	"github.com/dollarkillerx/light_rpc_benchmark/light_rpc/models"
)

func main() {
	client := client.NewClient(discovery.NewSimplePeerToPeer("127.0.0.1:8087", transport.TCP), client.SetPoolSize(100))
	connect, err := client.NewConnect("Server")
	if err != nil {
		log.Fatalln(err)
	}

	var response models.BenchmarkMessage
	err = connect.Call(light.DefaultCtx(), "Say", &models.BenchmarkMessage{Msg: "hello world"}, &response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
}
