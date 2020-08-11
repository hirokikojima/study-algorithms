# study-argorithms

## ベンチーマーク結果 (N = 100000)
```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/hirokikojima/study-algorithms/sort
BenchmarkSelectionSort-8               1        13808224798 ns/op            648 B/op          3 allocs/op
BenchmarkBubbleSort-8                  1        16604560034 ns/op              0 B/op          0 allocs/op
BenchmarkInsertionSort-8               1        1939112466 ns/op               0 B/op          0 allocs/op
BenchmarkHeapSort-8                  109          10706610 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/hirokikojima/study-algorithms/sort   34.305s
```