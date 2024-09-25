package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/

// 题目分析：可以买卖k次，根据之前的定义 dp[0][0]表示啥也不干， dp[0][1]表示第一次持有， dp[0][2]表示第二次不持有，
// dp[0][3]表示第二次持有，dp[0][4]表示第二次不持有，dp[0][5]表示第三次持有，dp[0][6]表示第3次持有
// 明确定义：dp[0][0]表示啥也不操作，dp[0][1]表示第一次持有股票最大利润，dp[0][2]表示第一次不持有股票最大利润，dp[0][3]表示第二次持有
// 股票最大利润，dp[0][4]表示第二次不持有股票最大利润，dp[0][5]表示第三次持有股票最大利润，dp[0][6]表示第三次不持有股票最大利润
// 递推公式：
// dp[i][0] = dp[i-1][0]
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
// dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i]])
// dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i]])
// dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i]])
// dp[i][5] = max(dp[i-1][5], dp[i-1][4]+prices[i]])
// dp[i][6] = max(dp[i-1][6], dp[i-1][5]+prices[i]])
// 初始化：dp[0][0]=0, dp[0][1]=-prices[i], dp[0][2]=0, dp[0][3]=-prices[0], dp[0][4]=0, dp[0][5]=-prices[i], dp[0][6]=0
func maxProfitIV02(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 0; i < k; i++ {
		// 2*k + 1 // 1, 3, 5
		// 2*k + 2 // 2, 4, 6
		dp[0][2*i+1] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 0; j < k; j++ {
			dp[i][2*j+1] = max(dp[i-1][2*j+1], dp[i-1][2*j]-prices[i])
			dp[i][2*j+2] = max(dp[i-1][2*j+2], dp[i-1][2*j+1]+prices[i])
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
		get := maxProfitIV02(tt.k, tt.prices)
		if get != tt.want {
			t.Fatalf("prices:%v, k%v, want:%v, get:%v", tt.prices, tt.k, tt.want, get)
		}
	}
}
