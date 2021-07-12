package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/async_utils"
	"github.com/dollarkillerx/light"
	"github.com/dollarkillerx/light/cryptology"
	"github.com/dollarkillerx/light_rpc_benchmark/grpc/proto"
	"github.com/montanaflynn/stats"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
)

var clientPool = make(chan proto.HelloClient, poolSize)
var poolSize = 100
var key = []byte("58a95a8f804b49e686f651a0d3f6e631")

func main() {
	for i := 0; i < poolSize; i++ {
		conn, e := grpc.Dial(":9001", grpc.WithInsecure()) // grpc.WithInsecure() 不安全的传输
		if e != nil {
			panic(e.Error())
		}
		client := proto.NewHelloClient(conn) // 注册上去
		clientPool <- client
	}

	over := make(chan struct{})
	poolFunc := async_utils.NewPoolFunc(poolSize, func() {
		close(over)
	})

	total := 100000

	summary := make([]int64, total, total)

	suResp := atomic.Uint64{}
	suOK := atomic.Uint64{}

	stTime := time.Now().UnixNano()

	for i := 0; i < total; i++ {
		idx := i

		poolFunc.Send(func() {
			client := <-clientPool
			defer func() {
				clientPool <- client
			}()

			n := time.Now().UnixNano()
			ctx := light.DefaultCtx()
			ctx.SetTimeout(time.Second * 6)
			// light rpc 不支持明文传输 所以加上
			msg := "hello world"
			resp, err := client.Say(context.TODO(), &proto.BenchmarkMessage{
				Msg: coding([]byte(msg)),
			})

			if err == nil {
				suResp.Add(1)

				r := decoding(resp.Rp)
				decoding(resp.Msg)

				if string(r) == "ok" {
					suOK.Add(1)
				}
			}

			r := time.Now().UnixNano() - n
			summary[idx] = r
		})
	}

	poolFunc.Over()
	fmt.Println("Send Over")
	<-over

	end := time.Now().UnixNano() - stTime
	endTime := end / 1000000
	fmt.Println("总耗时: ", endTime)
	totalF6 := make([]float64, 0, total)
	for _, k := range summary {
		totalF6 = append(totalF6, float64(k))
	}

	mean, _ := stats.Mean(totalF6)
	median, _ := stats.Median(totalF6)
	max, _ := stats.Max(totalF6)
	min, _ := stats.Min(totalF6)
	p99, _ := stats.Percentile(totalF6, 99.9)

	log.Printf("sent     requests    : %d\n", total)
	log.Printf("received requests    : %d\n", suResp.Load())
	log.Printf("received requests_OK : %d\n", suOK.Load())
	log.Printf("throughput  (TPS)    : %d\n", total*1000/int(endTime))
	log.Printf("mean: %.f ns, median: %.f ns, max: %.f ns, min: %.f ns, p99: %.f ns\n", mean, median, max, min, p99)
	log.Printf("mean: %d ms, median: %d ms, max: %d ms, min: %d ms, p99: %d ms\n", int64(mean/1000000), int64(median/1000000), int64(max/1000000), int64(min/1000000), int64(p99/1000000))
}

func coding(r []byte) []byte {
	encrypt, err := cryptology.AESEncrypt(key, r)
	if err != nil {
		log.Fatalln(err)
	}
	//sn, _ := codes.CompressorManager.Get(codes.Snappy)
	//zip, err := sn.Zip(encrypt)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	return encrypt
}

func decoding(r []byte) []byte {
	//sn, _ := codes.CompressorManager.Get(codes.Snappy)
	//rc, err := sn.Unzip(r)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	rb, err := cryptology.AESDecrypt(key, r)
	if err != nil {
		log.Fatalln(err)
	}

	return rb
}
