package _1_array

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/subsets-ii/

func subsetsWithDup(nums []int) [][]int {
	var backtracking func(nums []int, startIdx int)

	var res [][]int
	var path []int
	backtracking = func(nums []int, startIdx int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := startIdx; i < len(nums); i++ {
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}

	slices.Sort(nums)
	backtracking(nums, 0)
	return res
}

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{2, 1, 2}))
}
