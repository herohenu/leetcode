package leet

import "fmt"

//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array

func Test21() {
	nums := []int{1, 1, 1, 2, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	slow := 1
	for j := 1; j < len(nums); j++ {
		if nums[j-1] != nums[j] {
			nums[slow] = nums[j] // slow <- fast
			slow++

		}
		fmt.Println("step ", j, " nums : ", nums)
	}
	nums = nums[:slow]
	fmt.Println(nums)
	return slow
}
