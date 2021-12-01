package main

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	//遍历排序数组；
	//外层循环作为 第一个元素；
	//内层循环  和末尾倒着 求和 计算和 看能否为0

	//从小到大遍历这个数组，每次取一个元素，将这个元素的相反数设为target ok
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {

		target := 0 - nums[i]

		j := i + 1
		k := len(nums) - 1
		for j < k {

			if nums[j]+nums[k] == target {
				fmt.Println("solution element is : ", []int{nums[i], nums[j], nums[k]})
				res = append(res, []int{nums[i], nums[j], nums[k]})
				break
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}

	}

	return res
}

func ThreeSumMain() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	fmt.Println(nums)
	res := ThreeSum(nums)
	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	if len(nums) == 2 {
		if nums[0]+nums[1] == target {
			return []int{0, 1}
		} else {
			return []int{}
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {

				return []int{i, j}
			}
		}

	}

	return []int{}
}
