package main

import (
	"fmt"
	"math"
)

/*
@Time : 2021/11/8 9:36
*/
var dict = make(map[int]int)

func main() {
	var coins = []int{1, 2, 5}
	num := 13
	//t0 := time.Now().UnixNano()
	//result := dp(coins, num)
	//t1 := time.Now().UnixNano()
	//fmt.Println("result1 : ", result)
	//fmt.Println("result1 cost: ", t1-t0)
	//
	//result = dpWithDict(coins, num)
	//fmt.Println("result2 : ", result)
	//
	//t2 := time.Now().UnixNano()
	//
	//fmt.Println("dict : ", dict)
	//fmt.Println("result2 cost: ", t2-t1)

	result3 := getCoinChange(coins, num)
	fmt.Println("result3 : ", result3)

}

// return 返回的硬币个数
func dp(coins []int, n int) int {

	//fmt.Println(" will exec n is :", n)
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	res := math.MaxInt32

	for _, coin := range coins {
		subProblem := dp(coins, n-coin)
		//fmt.Printf(" coin is %d  subproblem(dp(%d)) result is : %d \n", coin, n-coin, subProblem)
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
		fmt.Printf(" use coin  %d , res is : %d \n ", coin, res)
	}

	if res == math.MaxInt32 {
		return -1
	}

	return int(res)

}

// 递归带字典解法
func dpWithDict(coins []int, n int) int {
	//fmt.Println(" will exec n is :", n)
	if n == 0 {

		return 0
	}
	if n < 0 {
		return -1
	}
	res := math.MaxInt32

	for _, coin := range coins {
		subProblem := dpWithDict(coins, n-coin)
		//fmt.Printf(" coin is %d  subproblem(dp(%d)) result is : %d \n", coin, n-coin, subProblem)
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
		fmt.Printf(" current %d ,use coin  %d , res is : %d \n ", n, coin, res)
	}

	if res == math.MaxInt32 {
		return -1
	}
	// 当所求的结果为n 时候 ，记录对应的res
	dict[n] = res

	return int(res)

}

// 迭代法
func getCoinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	res[0] = 0
	// value 初始化
	for i := 1; i <= amount; i++ {
		res[i] = math.MaxInt32
	}
	for i := 1; i < len(res); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			fmt.Printf("i =%d , res[%d]=%d \n", i, i, res[i])
			res[i] = min(res[i], 1+res[i-coin])
			fmt.Printf("set res[%d]= result %d \n  ", i, res[i])
		}

	}

	if res[amount] == amount+1 {
		return -1
	}

	return res[amount]

}

// compare which is samll
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
