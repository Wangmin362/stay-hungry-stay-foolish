package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/coin-change-ii/description/

func change(amount int, coins []int) int {
	// dp[j]为容量为j的背包，恰好可以装满的最大组合数
	// dp[j] += dp[j - nums[i]]
	dp := make([]int, amount+1)
	dp[0] = 1

	fmt.Println(dp)
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
		fmt.Println(dp)
	}

	return dp[amount]
}

func TestChange(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
}
