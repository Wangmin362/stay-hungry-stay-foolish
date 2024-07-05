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

		cache := map[int]int{}
		for idx := startIdx; idx < len(nums); idx++ {
			if cnt := cache[nums[idx]]; cnt > 0 { // 去重
				continue
			}

			if len(path) > 0 && nums[idx] < path[len(path)-1] {
				continue
			}

			cache[nums[idx]]++
			path = append(path, nums[idx])
			backtracking(nums, idx+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0)
	return res
}

func TestFindSubsequences(t *testing.T) {
	fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}))
}
