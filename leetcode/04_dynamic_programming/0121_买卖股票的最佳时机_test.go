package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	// dp[i][0] 表示第i天持有股票最多的现金价值 第i天持有股票的现金价值其实要么第i-1天已经持有了，因此未dp[i-1][0]
	// 或者是第i天把股票买进来，那么此时的现金价值就是-prices[i], 多以dp[i][0] = max(dp[i-1][0], -prices[i])
	// dp[i][1]表示第i天卖出股票的现金价值，因此第i天股票的价值可以是i-1天继续持有的价值，也即是dp[i-1][1]，也可以是
	// 第i天卖出股票的价值，也就是dp[i-1][0] + price[i]
	dp := make([][2]int, len(prices))
	dp[0] = [2]int{}
	dp[0][0] = -prices[0] // 第0天持有股票的价值，其实就是第0天买入股票的价值
	dp[0][1] = 0          // 第0天卖出股票的价值，因为没有买入，所以然时0
	for i := 1; i < len(prices); i++ {
		dp[i] = [2]int{}
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(-prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(prices[i]+dp[i-1][0])))
	}

	return dp[len(prices)-1][1]
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
