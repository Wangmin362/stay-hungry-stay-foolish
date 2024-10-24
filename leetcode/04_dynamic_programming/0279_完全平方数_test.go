package _0_basic

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/perfect-squares/description/

// 题目分析：背包的容量为n, 物品为[1, sqrt(n)]的平方，每个物品有无限个，求能够装满容量为n的最少数量，由于物品有1，所以任何容量的背包都可以装满
// 由于物品有无限个，因此是一个完全背包，所以背包顺序是从小到大。由于这里没有求组合或者排列，因此先背包还是先物品都可以，这里还是先物品，在背包
// 明确定义：dp[j]表示容量为j的背包，从[1, sqrt(j)]的平方数中去除最少的数目可以装满j
// 递推公式：dp[j] = min(dp[j- num] + 1, dp[j]), num是一个完全平方数字，从[1,sqrt(j)]中取一个
// 初始化：dp[0] = 0  dp[j] = math.MathInt
// dp大小：n+1
// 返回值：dp[n]
func numSquares0912(n int) int {
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt
	}

	// 先物品，在背包
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	// 先背包，在物品
	//for j := 0; j <= n; j++ {
	//	thres := int(math.Sqrt(float64(j)))
	//	for i := 1; i <= thres; i++ {
	//		if j >= i*i {
	//			dp[j] = min(dp[j], dp[j-i*i]+1)
	//		}
	//	}
	//}

	return dp[n]
}

// 递归：dfs(c) = min(dfs(c-i*i)+1)
func numSquaresDfs(n int) int {
	var dfs func(c int) int
	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = -1
	}

	dfs = func(c int) int {
		if c <= 1 {
			return c
		}
		if mem[c] != -1 {
			return mem[c]
		}

		res := math.MaxInt
		for i := 1; i*i <= c; i++ {
			res = min(res, dfs(c-i*i)+1)
		}
		mem[c] = res
		return res
	}

	res := dfs(n)
	return res
}

// 递归：dfs(c) = min(dfs(c-i*i)+1)
// 递推：f[c] = min(f[c-i*i]+1)
func numSquaresDp(n int) int {
	f := make([]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = math.MaxInt
	}
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if f[i-j*j] != math.MaxInt {
				f[i] = min(f[i], f[i-j*j]+1)
			}
		}
	}

	return f[n]
}

func TestNumSquares(t *testing.T) {
	var testdata = []struct {
		n    int
		want int
	}{
		{n: 12, want: 3},
		{n: 13, want: 2},
	}
	for _, tt := range testdata {
		get := numSquaresDp(tt.n)
		if get != tt.want {
			t.Fatalf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
