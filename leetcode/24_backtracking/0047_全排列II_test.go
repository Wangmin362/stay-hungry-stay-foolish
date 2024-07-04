package _1_array

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/permutations-ii/description/

func permuteUnique(nums []int) [][]int {
	var backtracking func(nums []int)
	var res [][]int
	var path []int

	cache := map[int]int{}
	for _, n := range nums { // 统计每个数字出现的次数
		cache[n]++
	}
	backtracking = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			if cnt := cache[nums[i]]; cnt <= 0 { // 说明这个重复数字不能再重复使用
				continue
			}
			path = append(path, nums[i])
			cache[nums[i]]--
			backtracking(nums)
			path = path[:len(path)-1]
			cache[nums[i]]++
		}
	}

	slices.Sort(nums)
	backtracking(nums)
	return res
}

func TestPrmuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
