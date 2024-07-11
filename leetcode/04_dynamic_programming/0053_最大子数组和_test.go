package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/maximum-subarray/description/

func maxSubArray(nums []int) int {
	// dp[i]定义为nums数组前i个元素，包括第i个元素的和
	// dp[i] = max(dp[i-1]+nums[i], nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
