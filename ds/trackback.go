package main

import "container/list"

/*
@Time : 2021/12/3 9:51
*/

var res = new(list.List)

func backtrack(nums []int, track map[int]bool) {
	if len(track) == len(nums) {

		//res.InsertAfter(res)

	}

	for i := 0; i < len(nums); i++ {
		// if nums[i]已经做过选择 跳过
		if track[nums[i]] {
			continue
		}
		// 做出选择
		track[nums[i]] = true

		// 进入下一层决策树
		backtrack(nums, track)

		//取消选择
		delete(track, nums[i])

	}

	//

}

func permute(nums []int) {
	track := make(map[int]bool)
	backtrack(nums, track)
	//return res
}
