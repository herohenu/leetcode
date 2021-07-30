package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@Time : 2021/7/26 14:11
*/

var wg sync.WaitGroup

func main() {
	chanMore()
}

func chanSlice() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		go func(ch chan int) {
			ch <- i
		}(chs[i])
	}
	//close(ch)
	for ch := range chs {
		fmt.Printf(" result : %v \n ", ch)
	}
	fmt.Println("end ... ")
}

// dead lock
func chanMore() {
	chs := make(chan int, 3)
	for i := 0; i < 10; i++ {
		//go func(i int) {
		wg.Add(1)
		chs <- i
		go Read(chs, i)
		//}(i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func Read(ch chan int, i int) {
	fmt.Printf("read i%d --- ch:  %+v time : %+v \n ", i, <-ch, time.Now())
	//<-ch
	wg.Done()

}
