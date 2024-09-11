package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/climbing-stairs/description/

// 题目分析：如果想要达到第n个台阶，那么一定是n-1台阶跨一步，以及n-2台阶跨两步过来的
// 明确DP定义 dp[i]表示想要到达第i个台阶不同的方法
// 状态转移公式： dp[i] = dp[i-1] + dp[i-2]
// 初始化： dp[0]=0, dp[1]=1, dp[2]=2
// 遍历顺序：从前往后，因为后面的状态是由前面的状态递推出来的
// dp数组大小: 需要计算[0, n]个状态，因此dp数组的长度为n+1
// 返回值： dp[n]

func climbStairs01(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 从递推公式可以知道，其实就是斐波那契数，第i个状态，只需要知道i-1以及i-2的状态，i-3以及之前的状态不需要记录，因此可以优化dp数组为O(1)的
// 空间复杂度
func climbStairs02(n int) int {
	if n <= 2 {
		return n
	}
	n1, n2 := 1, 2

	// n1, n2, dp
	//     n1, n2, dp
	for i := 3; i <= n; i++ {
		dpi := n1 + n2
		n1 = n2
		n2 = dpi
	}

	return n2
}

func TestClimbStairs(t *testing.T) {
	var testData = []struct {
		n    int
		want int
	}{
		{n: 2, want: 2},
		{n: 3, want: 3},
	}

	for _, tt := range testData {
		get := climbStairs01(tt.n)
		if get != tt.want {
			t.Errorf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
