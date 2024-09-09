package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sum-of-all-subset-xor-totals/description/?envType=problem-list-v2&envId=backtracking&difficulty=EASY

func subsetXORSum(nums []int) int {
	var backtracking func(start, xor int)

	var res int
	backtracking = func(start, xor int) {
		res += xor

		for i := start; i < len(nums); i++ {
			backtracking(i+1, xor^nums[i])
		}
	}

	backtracking(0, 0)
	return res
}

func TestSubsetXORSum(t *testing.T) {
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
}
