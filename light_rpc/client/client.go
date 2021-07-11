package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/async_utils"
	"github.com/dollarkillerx/light"
	"github.com/dollarkillerx/light/client"
	"github.com/dollarkillerx/light/discovery"
	"github.com/dollarkillerx/light/transport"
	"github.com/dollarkillerx/light_rpc_benchmark/light_rpc/models"
	"github.com/montanaflynn/stats"
	"go.uber.org/atomic"
)

func main() {
	client := client.NewClient(discovery.NewSimplePeerToPeer("127.0.0.1:8087", transport.TCP), client.SetPoolSize(100))
	connect, err := client.NewConnect("Server")
	if err != nil {
		log.Fatalln(err)
	}

	over := make(chan struct{})
	poolFunc := async_utils.NewPoolFunc(100, func() {
		close(over)
	})

	total := 1000000

	summary := make([]int64, total, total)

	suResp := atomic.Uint64{}
	suOK := atomic.Uint64{}

	stTime := time.Now().UnixNano()

	for i := 0; i < total; i++ {
		idx := i

		poolFunc.Send(func() {
			var response models.BenchmarkMessage
			n := time.Now().UnixNano()
			ctx := light.DefaultCtx()
			ctx.SetTimeout(time.Second * 6)
			err = connect.Call(ctx, "Say", &models.BenchmarkMessage{Msg: fmt.Sprintf("hello world :%d", idx)}, &response)
			r := time.Now().UnixNano() - n
			if err == nil {
				suResp.Add(1)
			}
			if response.Rp == "ok" {
				suOK.Add(1)
			}
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
