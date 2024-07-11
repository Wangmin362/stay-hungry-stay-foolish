package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLIS(nums []int) int {
	// dp[i]定义为数组中的前i个数字，包括数字i结合的最长递增子序列的长度
	// dp[i] = max(dp[i], dp[i-j] + 1)
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 至少长度是1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { // 大于才是递增子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
