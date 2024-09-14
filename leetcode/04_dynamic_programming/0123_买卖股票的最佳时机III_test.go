package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

/*
题目分析：只能买卖2两次股票，我们可以使用5个状态来表示，分别是：
0: 啥也不操作，继续保持之前的状态的最大的现金价值  1：表达第一次持有股票的现金价值  2：表示第一次不持有股票的现金价值
3：表示第二次持有股票的现金价值  4:表示第二次不持有股票的现金价值

明确定义：dp[i][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4]
dp[i][0] = dp[i-1][0]，啥也不操作的现金价值就是前一天的现金价值
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i]),和之前一样要么延续之前的持有状态，要么之前没有持有，第i天买入股票
dp[i][2] = max(dp[i-1][2], dp[i-1][1] + prices[i]),和之前一样，要么是延续之前不持有的状态，要么是之前已经持有，今天卖出
dp[i][3] = max(dp[i-1][3], dp[i-1][2] - prices[i]),要么是之前第二次持有股票的延续，要么是之前持有第一次股票已经卖出，今天买入
dp[i][4] = max(dp[i-1][4], dp[i-1][3] + prices[i]),要么是延续之前第二次持有股票的状态，要么是把之前第二次持有的股票卖出

初始化： dp[0][0] = 0, dp[0][1] = -prices[i], dp[0][2] = 0, dp[0][3] = -prices[i], dp[0][4] = 0
*/
func maxProfitIII(prices []int) int {
	dp := make([][5]int, len(prices))
	dp[0][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4] = 0, -prices[0], 0, -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]
}

// 状态压缩，进行优化
func maxProfitIII02(prices []int) int {
	zeo, one, two, thr, fou := 0, -prices[0], 0, -prices[0], 0

	for i := 1; i < len(prices); i++ {
		nzeo := zeo
		none := max(one, zeo-prices[i])
		ntwo := max(two, one+prices[i])
		nthr := max(thr, two-prices[i])
		nfou := max(fou, thr+prices[i])

		zeo = nzeo
		one = none
		two = ntwo
		thr = nthr
		fou = nfou
	}

	return fou
}

func TestMaxProfitIII(t *testing.T) {
	var testdata = []struct {
		prices []int
		want   int
	}{
		{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, want: 6},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
		{prices: []int{1, 2, 3, 4, 5}, want: 4},
	}
	for _, tt := range testdata {
		get := maxProfitIII02(tt.prices)
		if get != tt.want {
			t.Fatalf("prices:%v, want:%v, get:%v", tt.prices, tt.want, get)
		}
	}
}
