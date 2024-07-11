package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

func maxProfitII(prices []int) int {
	// dp[i][0]表示第i天持有股票的价值，第i天持有股票的价值，要么是i-1天已经持有了，因此dp[i-1][0]，要么是当天买入的价值，由于第i天需要买入，
	// 因此第i-1天必须卖出，因此当前的现金价值就是i-1天卖出股票所赚的钱减去今天买入的股票，也就是dp[i-1][1]-price[i]
	// dp[i][1]表示第i天卖出股票的价值，第i天卖出股票的价值，要么是i-1天已经买入了，也就是dp[i-1][1]，或者是第i天卖出的价值dp[i-1][0] + price[i]
	dp := make([][2]int, len(prices))
	dp[0] = [2]int{}
	dp[0][0] = -prices[0] // 第0天持有的价值，其实就是买入的价钱
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = [2]int{}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]-prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]+prices[i])))
	}

	return dp[len(prices)-1][1]
}

func TestMaxProfitII(t *testing.T) {
	fmt.Println(maxProfitII([]int{7, 1, 5, 3, 6, 4}))
}
