package main

import "math"

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//
//
//示例 1：
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func main3() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	println("max res:", res)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ele := nums[0]
	sum := 0
	for _, v := range nums {

		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		ele = int(math.Max(float64(sum), float64(ele)))
	}
	return sum
}
