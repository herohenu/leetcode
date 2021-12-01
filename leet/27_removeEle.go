package main

import "fmt"

func main2() {
	res := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println("res :", res)
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}

	nums = nums[:i]

	return i
}
