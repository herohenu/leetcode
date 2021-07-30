package main

import (
	"fmt"
	"time"
)

func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	ch <- fmt.Sprintf("Fib(%v):%v\n", number, x)
}

func fibsyn(number int) int {
	x, y := 1, 1
	for i := 0; i < number; i++ {
		x, y = y, x+y
	}
	return x
}

func FibStart() {
	start := time.Now()
	size := 30
	ch := make(chan string, size)
	for i := 0; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 0; i < size; i++ {
		fmt.Printf(" resutl :%s ", <-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds ! ", elapsed)
}

func FibSync() {
	start := time.Now()
	size := 30

	for i := 0; i < size; i++ {
		res := fibsyn(i)
		fmt.Printf("i =%d ,result: %v \n ", i, res)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds ! ", elapsed)
}
