package main

import (
	"fmt"

	"sync"
	"time"
)

/*
@Time : 2021/7/16 19:22
*/

func main() {
	//str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//i := 0
	//for {
	//
	//	j := i % len(str)
	//	fmt.Print("j : ", j)
	//	fmt.Println(" ", string(str[j]))
	//	i++
	//	time.Sleep(300 * time.Millisecond)
	//
	//}
	//work()
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	//fmt.Println("arr :", arr)
	//
	//for _, v := range arr {
	//	bad := leet.FirstBadVersion(v)
	//	fmt.Printf(" use val: %d ,  bad is :%d \n", v, bad)
	//}

	arr := []int{-4, -1, 0, 3, 10}
	SortedSquares(arr)

}

func binarySearch() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	for _, val := range nums {
		target := 2
		findidx := Search(nums, target)
		fmt.Printf("varl %d idx is : %d  \n", val, findidx)
	}
	fmt.Printf(" done ..... ")
}

func work() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Println(i)
				i++
				letter <- true
				time.Sleep(100 * time.Millisecond)
				break
				//default:
				//	break
				//}
			}

		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			i := 0
			for {
				select {
				case <-letter:
					if i > len(str)-1 {
						fmt.Println("-------done letter ")
						wg.Done()
						return
					}

					i = i % len(str)
					fmt.Print("i :", i)
					fmt.Println(string(str[i]))
					i++
					number <- true
					time.Sleep(100 * time.Millisecond)
					break
				default:
					break
				}
			}

		}(&wg)
		number <- true
		wg.Wait()
		fmt.Println("work ... ")
	}()
}
