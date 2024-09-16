package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/

// 明确定义：dp[i][0]表示持有股票获取的最大利用，dp[i][1]表示不持有股票获取的最大利润
// 状态方程：
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i]) 即第i天持有股票的最大利润只有两个可能性，其一是延续之前持有的股票，
// 也就是dp[i-1][0]，或者是之前没有持有股票，第i天买入股票，由于买入股票的最大利润为负数，所以这里是相加当前的股票，相当于计算出来最大利润
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i] - fee); 第i天不持有股票的最大利润也有两个可能性，其一是延续之前不持有股票
// 的最大利润，其二是第i天卖出股票，此时有手续费，因此需要减去手续费
// 初始化dp[0][0] = -prices[0], dp[0][1] = 0
// 遍历顺序：从前往后
func maxProfit714(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][0])
}

func TestMaxProfit714(t *testing.T) {
	var testdata = []struct {
		prices []int
		fee    int
		want   int
	}{
		{prices: []int{1, 3, 2, 8, 4, 9}, fee: 2, want: 8},
		{prices: []int{1, 3, 7, 5, 10, 3}, fee: 3, want: 6},
	}
	for _, tt := range testdata {
		get := maxProfit714(tt.prices, tt.fee)
		if get != tt.want {
			t.Fatalf("prices:%v, fee:%v, want:%v, get:%v", tt.prices, tt.fee, tt.want, get)
		}
	}
}
