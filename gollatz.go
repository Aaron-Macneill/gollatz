package main

import "fmt"

func collatz(n int64) bool{
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
	var i int64
	for {
		fmt.Println(i,collatz(i))
		i++
	}
}
