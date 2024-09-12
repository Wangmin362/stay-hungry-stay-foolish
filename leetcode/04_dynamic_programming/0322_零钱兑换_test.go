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

func coinChange02(coins []int, amount int) int {
	// dp[j] = min(dp[j], dp[j-coins[i]+1])
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

// 明确定义：dp[j]为总金额为j的背包可以凑成的最少硬币个数
// 状态方程：dp[j] = min(dp[j-nums[i]]+1， dp[j])
// 初始化: dp[0] = 1
// 遍历顺序：先物品，在背包，背包从小到大，因为是无限个，并且和顺序没有关系
func coinChange0911(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
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
	var testdata = []struct {
		coins  []int
		amount int
		want   int
	}{
		{coins: []int{1, 2, 5}, amount: 11, want: 3},
	}
	for _, tt := range testdata {
		get := coinChange0911(tt.coins, tt.amount)
		if get != tt.want {
			t.Fatalf("coins:%v, amount:%v, want:%v, get:%v", tt.coins, tt.amount, tt.want, get)
		}
	}
}
