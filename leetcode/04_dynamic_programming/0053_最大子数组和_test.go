package _0_basic

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/maximum-subarray/description/

// 解法一：暴力
// 解法二：动态规划
// 解法三：贪心

func maxSubArray01(nums []int) int {
	sum := func(start, end int) int {
		var res int
		for start <= end {
			res += nums[start]
			start++
		}
		return res
	}
	res := math.MinInt32
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			res = max(res, sum(i, j))
		}
	}

	return res
}

// 明确定义 ：dp[i]表示前i个数的最大和，即nums[0:i]子数组的最大和
// 递推公式：dp[i] = max(dp[i-1] + nums[i], nums[i]), 即要么是前nums[0,i-1]数组的最大和加上当前数字nums[i]，要么是当前数字，如果之前的和
// 加上现在的数字比现在的数字还小，那么就从当前数字开始计数
// 初始化：dp[0] = nums[0]
// 遍历顺序，从1开始，从小到大
func maxSubArray02(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, dp[i])
	}

	return res
}

func TestMaxSubArray(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
	}
	for _, tt := range testdata {
		get := maxSubArray02(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
