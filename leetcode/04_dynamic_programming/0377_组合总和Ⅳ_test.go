package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/combination-sum-iv/description/

// 题目分析：容量为target的背包，从nums中取数，一共有多少种方式可以凑成target
// 明确定义：dp[j]为容量为j的背包可以凑成的数量
// 转移方程： dp[j] += dp[j-nums[i]]
// 初始化：dp[0] = 1
// 遍历顺序：先背包，再物品，背包从小到大。   因为这道题目是有序的，也就是求排列的数量，那么就是先背包后物品
func combinationSum0912(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j < nums[i] {
				continue
			}
			dp[j] += dp[j-nums[i]]
		}
		fmt.Println(dp)
	}

	return dp[target]
}

// 本质上是爬楼梯问题
// 递归：dfs(i) += dfs(i-nums[j])
func combinationSum4Dfs(nums []int, target int) int {
	var dfs func(i int) int
	mem := make([]int, target+1)
	for i := 0; i <= target; i++ {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}

		if mem[i] != -1 {
			return mem[i]
		}

		var res int
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				res += dfs(i - nums[j])
			}
		}
		mem[i] = res
		return res
	}
	return dfs(target)
}

// 本质上是爬楼梯问题
// 递归：dfs(i) += dfs(i-nums[j])
// 递推：f[i] += f[i-nums[j]]
func combinationSum4Dp(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				f[j] += f[j-nums[i]]
			}
		}

	}

	return f[target]
}

func TestCombinationSum4(t *testing.T) {
	var testdata = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 2, 3}, target: 4, want: 7},
	}

	for _, tt := range testdata {
		get := combinationSum4Dp(tt.nums, tt.target)
		if get != tt.want {
			t.Fatalf("nums:%v, target:%v, want:%v, get:%v", tt.nums, tt.target, tt.want, get)
		}
	}
}
