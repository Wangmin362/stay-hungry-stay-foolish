package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/

/*
题目分析：可以买卖k次，根据之前的定义 dp[0][0]表示啥也不干， dp[0][1]表示第一次持有， dp[0][2]表示第二次不持有，
dp[0][3]表示第二次持有，dp[0][4]表示第二次不持有，dp[0][5]表示第三次持有，dp[0][6]表示第3次持有


*/

func maxProfitIV(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}

	dp[0][0] = 0
	for j := 0; j < 2*k; j += 2 {
		dp[0][j+1] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}

	return dp[len(prices)-1][2*k]
}

func TestMaxProfitIV(t *testing.T) {
	var testdata = []struct {
		prices []int
		k      int
		want   int
	}{
		{prices: []int{2, 4, 1}, k: 2, want: 2},
		{prices: []int{3, 2, 6, 5, 0, 3}, k: 2, want: 7},
		{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, k: 2, want: 6},
	}
	for _, tt := range testdata {
		get := maxProfitIV(tt.k, tt.prices)
		if get != tt.want {
			t.Fatalf("prices:%v, k%v, want:%v, get:%v", tt.prices, tt.k, tt.want, get)
		}
	}
}
