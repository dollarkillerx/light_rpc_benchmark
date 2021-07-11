# light_rpc_benchmark
light rpc benchmark

light v0.0.1 bate 全加密
``` 
Send Over
总耗时:  124103
2021/07/11 23:21:39 sent     requests    : 1000000
2021/07/11 23:21:39 received requests    : 999313
2021/07/11 23:21:39 received requests_OK : 999313
2021/07/11 23:21:39 throughput  (TPS)    : 8057
2021/07/11 23:21:39 mean: 11650181 ns, median: 4832682 ns, max: 6018865638 ns, min: 139708 ns, p99: 58745012 ns
2021/07/11 23:21:39 mean: 11 ms, median: 4 ms, max: 6018 ms, min: 0 ms, p99: 58 ms
```

grpc 半加密
``` 
Send Over
总耗时:  85003
2021/07/11 23:14:28 sent     requests    : 1000000
2021/07/11 23:14:28 received requests    : 1000000
2021/07/11 23:14:28 received requests_OK : 1000000
2021/07/11 23:14:28 throughput  (TPS)    : 11764
2021/07/11 23:14:28 mean: 8434775 ns, median: 7205904 ns, max: 101241805 ns, min: 172806 ns, p99: 42037508 ns
2021/07/11 23:14:28 mean: 8 ms, median: 7 ms, max: 101 ms, min: 0 ms, p99: 42 ms
```

go version go1.16.4 linux/amd64
CPU: AMD Ryzen 7 5800X 8-Core Processor
MEM: DDR4 32G 2666MHz