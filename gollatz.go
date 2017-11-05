package main

func Collatz(n int64) bool{
	for n > 1 {
		if(n%2 == 0) {
			n = n/2
			//fmt.Println(n)
		}else{
			n = 3*n+1
			//fmt.Println(n)
		}
	}
	return true
}

/*
func benchmark(workers *int, maxInt int) {
	guard := make(chan struct{}, workers)
	for i := int64(1); i<maxInt; i++ {
		guard <- struct{}{}
		go func(n int64) {
			Collatz(n)
			<-guard
		}(i)


func init() {
	var workers = flag.Int("threads", 30, "Number of goroutines to use")
}
*/
