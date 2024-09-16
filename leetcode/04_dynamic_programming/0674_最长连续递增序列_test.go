package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/

// 解法一：滑动窗口
// 解法二：回溯
// 解法三：贪心
// 解法四：动态规划

func findLengthOfLCISBacktracking(nums []int) int {
	var backtracking func(start int)

	var res int
	var path []int
	backtracking = func(start int) {
		res = max(res, len(path))

		for i := start; i < len(nums); i++ {
			if len(path) > 0 && nums[path[len(path)-1]] >= nums[i] {
				continue
			}
			if len(path) > 0 && path[len(path)-1]+1 != i {
				continue
			}
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return res
}

// 题目分析：求出最长连续的递增子序列
// 明确定义：dp[i]表示数组下标[0,i]的数组的最长连续递增子序列
// 递推公式：if dp[j] < dp[i] { dp[i] =  dp[j] + 1 } 只需要比较i的前一个数字即可，因为是连续递增子序列
// 初始化： dp[i] = 1, 最小就是1
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
		res = max(res, dp[i])
	}

	return res
}

func TestFindLengthOfLCIS(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 3},
		{nums: []int{1, 3, 5, 4, 7}, want: 3},
	}

	for _, tt := range testdata {
		get := findLengthOfLCIS(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v get:%v", tt.nums, tt.want, get)
		}
	}
}
