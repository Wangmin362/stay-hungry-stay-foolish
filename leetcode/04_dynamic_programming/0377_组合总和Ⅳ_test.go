package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/combination-sum-iv/description/

// 题目转换：容量为target的背包，从nums当中选取不同的数字放入到背包中，求有多少种组合可以装满背包
// 这是一个排列问题，因为1,2,1和1,1,2是不同的组合
func combinationSum4(nums []int, target int) int {
	// dp[j]定义为容量为j的背包，从0..i个数字当中可以装满背包次数
	// dp[j] += dp[j - nums[i]]
	// 结果为dp[target]
	dp := make([]int, target+1)
	dp[0] = 1
	fmt.Println("背包容量为: ", 0, dp)
	for j := 1; j <= target; j++ { // 先遍历背包
		for i := 0; i < len(nums); i++ { // 在遍历物品
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
		fmt.Println("背包容量为: ", j, dp)
	}

	return dp[target]
}

func combinationSum6(nums []int, target int) int {
	// dp[j] += dp[j - nums[i])
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

func TestCombinationSum4(t *testing.T) {
	fmt.Println(combinationSum6([]int{1, 2, 3}, 4))
}
