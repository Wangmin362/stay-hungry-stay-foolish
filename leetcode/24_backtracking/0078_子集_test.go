package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) [][]int {
	var backtracking func(nums []int, k, startIdx int)

	var res [][]int
	res = append(res, []int{})
	var path []int
	backtracking = func(nums []int, k, startIdx int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for idx := startIdx; idx < len(nums); idx++ {
			if len(nums)-idx < k-len(path) {
				continue
			}
			path = append(path, nums[idx])
			backtracking(nums, k, idx+1)
			path = path[:len(path)-1]
		}
	}

	for i := 1; i <= len(nums); i++ {
		backtracking(nums, i, 0)
	}

	return res
}

func subsets02(nums []int) [][]int {
	var backtracking func(nums []int, startIdx int)

	var res [][]int
	var path []int
	backtracking = func(nums []int, startIdx int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for idx := startIdx; idx < len(nums); idx++ {
			path = append(path, nums[idx])
			backtracking(nums, idx+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0)

	return res
}

func TestSubsets(t *testing.T) {
	fmt.Println(subsets02([]int{1, 2, 3}))
}
