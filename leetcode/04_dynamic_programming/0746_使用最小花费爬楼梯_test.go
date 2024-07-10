package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/min-cost-climbing-stairs/description/

// 实际上就是要爬到第n个台阶
func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	// dp[n]定义为爬到第n个台阶的最小费用
	// dp[n] = max(dp[n-1]+cost[n-1], dp[n-2]+cost[n-2])
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for n := 2; n <= len(cost); n++ {
		dp[n] = min(dp[n-1]+cost[n-1], dp[n-2]+cost[n-2])
	}

	return dp[len(cost)]
}
func TestMinCostClimbingStairs(t *testing.T) {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
