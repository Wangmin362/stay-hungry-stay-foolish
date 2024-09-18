package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/min-cost-climbing-stairs/description/

// 题目分析：从台阶i可以向上跳一格，或者跳两格，但是需要支持cost[i]费用
// 明确dp定义： dp[i]表示跳到第i个台阶需要的最小花费
// 状态转移方程：dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
// 初始化：dp[0],dp[1] = 0
// 遍历顺序：从前往后
// dp数组大小：[0, len(cost)]，也就是len(cost)+1
// 返回值dp[len(const)]

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}

	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

// 优化
func minCostClimbingStairs01(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}

	n1, n2 := 0, 0
	// n1, n2, dpi
	//     n1, n2, dpi
	for i := 2; i <= len(cost); i++ {
		dpi := min(n1+cost[i-2], n2+cost[i-1])
		n1 = n2
		n2 = dpi
	}

	return n2
}

// 明确定义：dp[i]表示爬到第i个调节需要支付的最小费用
// 递推公式：dp[i] = min(dp[i-1] + cost[i-1], dp[n-2] + cost[i-2])
// 初始化：dp[0], dp[1] = 0, 0
// 遍历顺序：从小到大
// 数组大小：n+1，因为需要爬到第n个台阶
func minCostClimbingStairs03(cost []int) int {
	dp0, dp1 := 0, 0
	// dp0, dp1, dp
	//      dp0, dp1, dp
	for i := 2; i <= len(cost); i++ {
		dp := min(dp1+cost[i-1], dp0+cost[i-2])
		dp0 = dp1
		dp1 = dp
	}
	return dp1
}

func TestMinCostClimbingStairs(t *testing.T) {
	var testData = []struct {
		cost []int
		want int
	}{
		{cost: []int{10, 15, 20}, want: 15},
		{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
	}

	for _, tt := range testData {
		get := minCostClimbingStairs03(tt.cost)
		if get != tt.want {
			t.Fatalf("cost:%v, want:%v, get:%v", tt.cost, tt.want, get)
		}
	}
}
