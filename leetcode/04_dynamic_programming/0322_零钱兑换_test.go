package _0_basic

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/coin-change/

// 题目分析：本题可以抽象为一个背包问题，容量为amount, 物品为coins，并且物品的价值就是重量。由于询问的是最少数量，因此没有顺序问题，所以
// 先物品在背包，由于每个物品是无限个，因此是一个完全背包问题，一次物品从小到大
// 明确定义：dp[j]表示容量为j的背包，可以由前i个硬币凑出来的最小数量
// 递推公式：dp[j] = min(dp[j], dp[j - coins[i]] + 1) // 要么选当前硬币，要么不选当前硬币
// 初始化：dp[0]:=0, dp[j] = math.MaxInt32
// 遍历顺序：先物品，在背包。 物品从小到大
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt { // 说明当前dp[j - coins[i]]并不能凑出来，直接跳过
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
		//fmt.Println(dp)
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
		//{coins: []int{1, 2, 5}, amount: 11, want: 3},
		//{coins: []int{3}, amount: 2, want: -1},
		{coins: []int{2}, amount: 3, want: -1},
	}
	for _, tt := range testdata {
		get := coinChange(tt.coins, tt.amount)
		if get != tt.want {
			t.Fatalf("coins:%v, amount:%v, want:%v, get:%v", tt.coins, tt.amount, tt.want, get)
		}
	}
}
