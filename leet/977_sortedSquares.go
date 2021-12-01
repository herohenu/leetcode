package main

import (
	"math"
	"sort"
)

/*
@Time : 2021/7/29 10:31
*/
//nums = [-4,-1,0,3,10]

func SortedSquares(nums []int) []int {

	for k, v := range nums {
		nums[k] = int(math.Pow(math.Abs(float64(v)), float64(2)))
	}

	sort.Ints(nums)
	return nums
}
