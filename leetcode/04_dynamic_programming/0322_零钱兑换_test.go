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

// 递归：dfs(i, c) = min(dfs(i-1, c), dfs(i, c-coins[i])+1 )
func coinChangeDfs(coins []int, amount int) int {
	var dfs func(i, c int) int
	mem := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		mem[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			mem[i][j] = -1
		}
	}

	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2 // 因为要比较最小值，所以要给一个大点的数
		}

		if mem[i][c] != -1 {
			return mem[i][c]
		}

		if c < coins[i] {
			res := dfs(i-1, c)
			mem[i][c] = res
			return res
		}
		res := min(dfs(i-1, c), dfs(i, c-coins[i])+1)
		mem[i][c] = res
		return res
	}
	if dfs(len(coins)-1, amount) >= math.MaxInt/2 {
		return -1
	}

	return dfs(len(coins)-1, amount)
}

// 递归：dfs(i, c) = min(dfs(i-1, c), dfs(i, c-coins[i])+1 )
// 递推：f[i][c] = min(f[i-1][c], f[i][c-coins[i]]+1 )
// 两边同时加一：可得：
// 递推：f[i+1][c] = min(f[i][c], f[i+1][c-coins[i]]+1 )
func coinChangeDp(coins []int, amount int) int {
	f := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		f[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			f[i][j] = math.MaxInt
		}
	}
	f[0][0] = 0

	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j < coins[i] {
				f[i+1][j] = f[i][j]
			} else {
				if f[i+1][j-coins[i]] == math.MaxInt {
					f[i+1][j] = f[i][j]
				} else {
					f[i+1][j] = min(f[i][j], f[i+1][j-coins[i]]+1)
				}
			}
		}
	}
	if f[len(coins)][amount] == math.MaxInt {
		return -1
	}

	return f[len(coins)][amount]
}

// 递归：dfs(i, c) = min(dfs(i-1, c), dfs(i, c-coins[i])+1 )
// 递推：f[i][c] = min(f[i-1][c], f[i][c-coins[i]]+1 )
// 两边同时加一：可得：
// 递推：f[i+1][c] = min(f[i][c], f[i+1][c-coins[i]]+1 )
// 优化为两行
// 递推：f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-coins[i]]+1 )
func coinChangeDp02(coins []int, amount int) int {
	f := make([][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			f[i][j] = math.MaxInt
		}
	}
	f[0][0] = 0

	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j < coins[i] {
				f[(i+1)%2][j] = f[i%2][j]
			} else {
				if f[(i+1)%2][j-coins[i]] == math.MaxInt {
					f[(i+1)%2][j] = f[i%2][j]
				} else {
					f[(i+1)%2][j] = min(f[i%2][j], f[(i+1)%2][j-coins[i]]+1)
				}
			}
		}
	}
	if f[len(coins)%2][amount] == math.MaxInt {
		return -1
	}

	return f[len(coins)%2][amount]
}

// 递归：dfs(i, c) = min(dfs(i-1, c), dfs(i, c-coins[i])+1 )
// 递推：f[i][c] = min(f[i-1][c], f[i][c-coins[i]]+1 )
// 两边同时加一：可得：
// 递推：f[i+1][c] = min(f[i][c], f[i+1][c-coins[i]]+1 )
// 优化为一行
// 递推：f[c] = min(f[c], f[c-coins[i]]+1 )
func coinChangeDp03(coins []int, amount int) int {
	f := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		f[j] = math.MaxInt
	}
	f[0] = 0

	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				if f[j-coins[i]] == math.MaxInt {
					f[j] = f[j]
				} else {
					f[j] = min(f[j], f[j-coins[i]]+1)
				}
			}
		}
	}
	if f[amount] == math.MaxInt {
		return -1
	}

	return f[amount]
}

func TestCoinChange(t *testing.T) {
	var testdata = []struct {
		coins  []int
		amount int
		want   int
	}{
		{coins: []int{1, 2, 5}, amount: 11, want: 3},
		{coins: []int{3}, amount: 2, want: -1},
		{coins: []int{2}, amount: 3, want: -1},
	}
	for _, tt := range testdata {
		get := coinChangeDp03(tt.coins, tt.amount)
		if get != tt.want {
			t.Fatalf("coins:%v, amount:%v, want:%v, get:%v", tt.coins, tt.amount, tt.want, get)
		}
	}
}
