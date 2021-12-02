package leet

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{-1, -1}
	}

	if len(nums) == 2 {
		if nums[0]+nums[1] == target {
			return []int{0, 1}
		} else {
			return []int{-1, -1}
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {

				return []int{i, j}
			}
		}

	}

	return []int{-1, -1}
}
