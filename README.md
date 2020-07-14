# study-argorithms
## Benchmark
```
search$ go test -bench . -benchmem
---
BenchmarkLinearSearch-8           346119              3311 ns/op               0 B/op          0 allocs/op
BenchmarkBinarySearch-8         54259042                21.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashSearch-8            6172159               189 ns/op               0 B/op          0 allocs/op
```
