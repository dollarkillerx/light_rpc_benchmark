# light_rpc_benchmark
light rpc benchmark

light v0.0.1 bate 全加密
``` 
Send Over
总耗时:  10356
2021/07/12 22:01:22 sent     requests    : 1000000
2021/07/12 22:01:22 received requests    : 1000000
2021/07/12 22:01:22 received requests_OK : 1000000
2021/07/12 22:01:22 throughput  (TPS)    : 96562
2021/07/12 22:01:22 mean: 988658 ns, median: 692178 ns, max: 28222823 ns, min: 41207 ns, p99: 16710486 ns
2021/07/12 22:01:22 mean: 0 ms, median: 0 ms, max: 28 ms, min: 0 ms, p99: 16 ms

Send Over
总耗时:  8974
2021/07/12 21:58:37 client.go:79: sent     requests    : 1000000
2021/07/12 21:58:37 client.go:80: received requests    : 1000000
2021/07/12 21:58:37 client.go:81: received requests_OK : 1000000
2021/07/12 21:58:37 client.go:82: throughput  (TPS)    : 111433
2021/07/12 21:58:37 client.go:83: mean: 822389 ns, median: 466363 ns, max: 34675799 ns, min: 37220 ns, p99: 14228576 ns
2021/07/12 21:58:37 client.go:84: mean: 0 ms, median: 0 ms, max: 34 ms, min: 0 ms, p99: 14 ms
```

grpc 
``` 
Send Over
总耗时:  1540
2021/07/12 21:52:48 sent     requests    : 100000
2021/07/12 21:52:48 received requests    : 100000
2021/07/12 21:52:48 received requests_OK : 100000
2021/07/12 21:52:48 throughput  (TPS)    : 64935
2021/07/12 21:52:48 mean: 1481574 ns, median: 1340954 ns, max: 6808181 ns, min: 71213 ns, p99: 5512718 ns
2021/07/12 21:52:48 mean: 1 ms, median: 1 ms, max: 6 ms, min: 0 ms, p99: 5 ms

Send Over
总耗时:  9169
2021/07/12 21:57:57 sent     requests    : 1000000
2021/07/12 21:57:57 received requests    : 1000000
2021/07/12 21:57:57 received requests_OK : 1000000
2021/07/12 21:57:57 throughput  (TPS)    : 109063
2021/07/12 21:57:57 mean: 869900 ns, median: 687608 ns, max: 21728321 ns, min: 61546 ns, p99: 5991032 ns
2021/07/12 21:57:57 mean: 0 ms, median: 0 ms, max: 21 ms, min: 0 ms, p99: 5 ms
```

go version go1.16.4 linux/amd64

CPU: AMD Ryzen 7 5800X 8-Core Processor

MEM: DDR4 32G 2666MHz