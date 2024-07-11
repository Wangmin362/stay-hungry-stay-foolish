package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/

func findLengthOfLCIS(nums []int) int {
	// dp[i]定义为nums数组前i个元素，包括第i个元素的连续最长组序列的长度
	// dp[i] =  dp[j]+1
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
