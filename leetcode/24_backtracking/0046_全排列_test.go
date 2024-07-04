package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/permutations/description/

func permute(nums []int) [][]int {
	var backtracking func(nums []int, cache map[int]struct{})

	var res [][]int
	var path []int
	backtracking = func(nums []int, cache map[int]struct{}) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := cache[nums[i]]; ok {
				continue
			}
			path = append(path, nums[i])
			cache[nums[i]] = struct{}{}
			backtracking(nums, cache)
			path = path[:len(path)-1]
			delete(cache, nums[i])
		}
	}

	backtracking(nums, map[int]struct{}{})
	return res
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{0, 1}))
}
