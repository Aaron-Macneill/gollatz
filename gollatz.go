package main

import (
	"fmt"
)
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


func main() {
	workers := 30
	guard := make(chan struct{}, workers)
	for i := int64(1); ; i++{
		guard <- struct{}{}
		go func(n int64) {
			fmt.Println(i, Collatz(n))
			<-guard
		}(i)
	}
}
