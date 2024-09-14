package _0_basic

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

// 直接暴力搜索，两层for循环，找到插值最大的差就是结果
func maxProfitN2(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

// 贪心算法，只需要保证左边取到最小值，右边取一个最大值即可
func maxProfitGreed(prices []int) int {
	res, low := 0, math.MaxInt32
	for i := 0; i < len(prices); i++ {
		low = min(low, prices[i])
		res = max(res, prices[i]-low)
	}
	return res
}

/*
题目分析：股票买入之前不能卖出，股票买入之后只能选择卖出，不能再次买入。同时，卖出之后也不能选择再次买入
也就是说，在某一天，我对于股票的行为只能是买入或者卖出。买入的前提是之前没有买入，卖出的前提是之前已经买入。

明确定义：dp[i][0]表示第i天持有股票的现金价值，dp[i][1]表示第i天不持有股票的现金价值

状态方程：dp[i][0]表示第i天持有股票现金价值，第i天持有股票有两种情况，要么是之前已经买入了，第i天延续之前的持有状态，此时dp[i][0] = dp[i-1][0]
要么是之前没有买入，第i天买入，此时dp[i][0] = -price[i]，此时现金价值就是负数，因为还没有卖出。综上dp[i][0] = max(dp[i-1][0], -price[i])
dp[i][1]表示第i天不持有股票的现金价值，第i天不持有股票，有两种情况，其一：之前就已经不持有股票了，只不过现在是之前的延续，这种情况下，
dp[i][1] = dp[i-1][1]，要么是之前已经持有股票，第i天选择卖出股票，此时也是不持有，此时又dp[i][1] = price[i]+dp[i-1][0]，由于之前
买入股票的现金价值为负数，因此这里直接相加即可算出中间获得的利益，综上所述：dp[i][1] = max(dp[i-1][1], price[i]+dp[i-1][1])
总和下来就是：dp[i][0] = max(dp[i-1][0], -price[i])  dp[i][1] = max(dp[i-1][1], price[i] + dp[i-1][0])

初始化：从公式可以看出，i应该从1开始计算，也就是dp[0][0],dp[0][1]进行初始化，dp[0][0]=-price[0]  dp[0][1]=0

dp数组大小: dp := make([][2]int, len(prices))

返回值：dp[len(prices)-1][1] 一定是最后一天不持有股票的价值，持有股票都没有意义了
*/
func maxProfitDp(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}

	return dp[len(prices)-1][1]
}

// 由于第i天只依赖于前一天的状态，因此可以压缩数组
func maxProfitDp02(prices []int) int {
	hold, nonHold := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		currHold := max(hold, -prices[i])
		currNonHold := max(nonHold, prices[i]+hold)

		hold = currHold
		nonHold = currNonHold
	}

	return nonHold
}

func TestMaxProfit(t *testing.T) {
	var testdata = []struct {
		prices []int
		want   int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{prices: []int{7, 6, 4, 3, 1}, want: 0},
	}
	for _, tt := range testdata {
		get := maxProfitDp02(tt.prices)
		if get != tt.want {
			t.Fatalf("prices:%v, want:%v, get:%v", tt.prices, tt.want, get)
		}
	}
}
