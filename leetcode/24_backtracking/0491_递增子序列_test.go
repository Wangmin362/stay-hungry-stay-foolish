package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/non-decreasing-subsequences/description/

func findSubsequences(nums []int) [][]int {
	var backtracking func(nums []int, startIdx int)

	var res [][]int
	var path []int
	backtracking = func(nums []int, startIdx int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for idx := startIdx; idx < len(nums); idx++ {
			if idx > startIdx && nums[idx] == nums[idx-1] {
				continue
			}
			path = append(path, nums[idx])
			backtracking(nums, idx+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0)
	return res
}

func TestFindSubsequences(t *testing.T) {
	fmt.Println(findSubsequences([]int{6, 4, 7, 7}))
}
