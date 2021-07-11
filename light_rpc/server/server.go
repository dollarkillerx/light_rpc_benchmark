package main

import (
	"fmt"
	"log"

	"github.com/dollarkillerx/light"
	"github.com/dollarkillerx/light/server"
	"github.com/dollarkillerx/light_rpc_benchmark/light_rpc/models"
)

func main() {
	ser := server.NewServer()
	err := ser.Register(&Server{})
	if err != nil {
		log.Fatalln(err)
	}

	err = ser.Run(server.UseTCP("0.0.0.0:8087"))
	if err != nil {
		log.Fatalln(err)
	}
}

type Server struct{}

func (s *Server) Say(ctx *light.Context, request *models.BenchmarkMessage, response *models.BenchmarkMessage) error {
	response.Msg = fmt.Sprintf("resp: %s", request.Msg)
	response.Rp = "ok"
	return nil
}
