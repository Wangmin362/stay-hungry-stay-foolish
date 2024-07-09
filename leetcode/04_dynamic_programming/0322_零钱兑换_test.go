package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/coin-change/

func coinChange(coins []int, amount int) int {
	// dp[j]定义为0..i个硬币可以转满背包容量为j的最少的硬币个数
	// dp[j] = min(dp[j], dp[j-coins[i]]+1)
	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i]])+1))
			}
		}
		fmt.Println(dp)
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}
