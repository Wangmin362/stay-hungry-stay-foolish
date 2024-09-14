package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

/*
题目分析：可以买卖多次，只不过每次买入之前的状态必须是没有持有股票的状态
明确定义：dp[i][0]表示第i天持有股票的价值，  dp[i][1]表示第i天不持有股票的价值

状态转移：第i天持有股票的价值有这么几种情况：其一：之前已经持有，也就是说dp[i][0] = dp[i-1][0]。 第二：之前没有持有股票，第i天持有股票，此时
dp[i][0] = dp[i-1][0] - prices[i]，综上dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
第i天不持有股票的价值还是和之前一样， dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])

初始化：dp[0][0]=-prices[0], dp[0][1] = 0
*/
func maxProfitII(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i][0]+prices[i])
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

// 优化，状态压缩，当前的状态只依赖于前一个状态
func maxProfitII02(prices []int) int {
	hold, nonHold := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		currHold := max(hold, nonHold-prices[i])
		currNonHold := max(nonHold, hold+prices[i])

		hold = currHold
		nonHold = currNonHold
	}

	return max(hold, nonHold)
}

func TestMaxProfitII(t *testing.T) {
	var testdata = []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 7},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
		{prices: []int{1, 2, 3, 4, 5}, want: 4},
	}
	for _, tt := range testdata {
		get := maxProfitII02(tt.prices)
		if get != tt.want {
			t.Fatalf("prices:%v, want:%v, get:%v", tt.prices, tt.want, get)
		}
	}
}
