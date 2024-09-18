package _1_array

import (
	"math"
	"testing"
)

func canJumpII(nums []int) int {
	if nums[0] == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	iMax := 0 // 初始化能跳到最远的位置就是0
	idx, jump := 0, 0
	// 遍历数组，如果当前最远的在当前位置的后面，说明当前位置时可以跳到的，因此
	// 只要当前位置加上可以跳到的最远位置大于记录的最远位置，就更新最远位置，最后比较
	// 最远位置和最后一个位置
	for idx, jump = range nums {
		if iMax >= idx && idx+jump > iMax {
			iMax = idx + jump
		}
		if iMax >= len(nums)-1 { // 只要能跳的最大位置超过最后的位置，那么就可以达到，此时直接返回
			return idx + 1
		}
	}

	return idx + 1
}

// 动态规划
// 明确定义：dp[i]表示跳到下标为i需要的最少的次数
// 递推公式：dp[i] = min(dp[i-1]+1, dp[i-2]+i, dp[i-3]+1), 当然，前提是nums[i-1]可以跳到i的位置
// 初始化：dp[0] = 0
// 遍历顺序：从小到达
func canJumpII0918(nums []int) int {
	dp := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		minStep := math.MaxInt32
		for j := 0; j < i; j++ {
			if j+nums[j] >= i { // 说明从j位置可以跳到i位置
				minStep = min(minStep, dp[j]+1)
			}
		}
		if minStep == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = minStep
		}
	}

	return dp[len(nums)-1]
}

func TestCanJumpII(t *testing.T) {
	twoSumTest := []struct {
		array  []int
		expect int
	}{
		//{array: []int{1, 2, 1, 1, 1}, expect: 3},
		//{array: []int{2, 3, 1, 1, 4}, expect: 2},
		{array: []int{3, 2, 1, 0, 4}, expect: -1},
	}

	for _, test := range twoSumTest {
		get := canJumpII0918(test.array)
		if get != test.expect {
			t.Errorf("target:%v, expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
