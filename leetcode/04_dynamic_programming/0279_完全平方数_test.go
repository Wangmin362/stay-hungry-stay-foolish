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

func TestNumSquares(t *testing.T) {
	var testdata = []struct {
		n    int
		want int
	}{
		{n: 12, want: 3},
		{n: 13, want: 2},
	}
	for _, tt := range testdata {
		get := numSquares0912(tt.n)
		if get != tt.want {
			t.Fatalf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
