package main

import (
	"flag"
)

func Benchmark(workers int, maxInt int64) {
	guard := make(chan struct{}, workers)
	for i := int64(1); i<maxInt; i++ {
		guard <- struct{}{}
		go func(n int64) {
			Collatz(n)
			<-guard
		}(i)
	}
}

func main() {
	var maxInt = flag.Int64("upto", 1e6, "Highest integer to collatz")
	var workers = flag.Int("groutes", 30, "Number of goroutines to use")
	flag.Parse()
	Benchmark(*workers, *maxInt)

}
