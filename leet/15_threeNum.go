package leet

import (
	"fmt"
	"sort"
)

//
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	m := make(map[string]bool)

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		//fmt.Println("now  i is : ", i)
		j := i + 1
		k := len(nums) - 1
		//fmt.Printf(" now idx   j %d , k %d \n", j, k)
		for j < k {
			if nums[j]+nums[k] == target {
				//fmt.Println("solution element is : ", []int{nums[i], nums[j], nums[k]})
				ijk := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
				//fmt.Println("ijk is :", ijk)
				if _, ok := m[ijk]; !ok {
					m[ijk] = true
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				j++
				k--
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}

	}

	return res
}

func TestThreeSum() {
	nums := []int{-2, 0, 1, 1, 2}
	nums = []int{-4, -2, -1, -1, 0, 1, 1, 2}
	// [[-2,0,2],[-2,1,1]]
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
