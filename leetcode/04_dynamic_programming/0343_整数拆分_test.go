package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/integer-break/description/

// 题目分析：一个证书n，可以拆分为(1, n-1), (2, n-2), (3, n-3), (4, n-4), (5, n-5), ... (n-1, 1)
// 那么 dp[n] = max(dp[n-1]*1, dp[n-2]*2, dp[n-3]*3, dp[n-4]*4, ...dp[2]*(n-2), dp[1]*(n-1))
// 明确定义：dp[n]为数字n可以拆分为整数的最大值
// 转移方程：dp[n] = max(dp[n-1]*1, dp[n-2]*2, dp[n-3]*3, dp[n-4]*4, ...dp[2]*(n-2), dp[1]*(n-1))
// 初始化：dp[0]=0, dp[1]=1, dp[2]=1, dp[3]=2
// 遍历顺序：从前往后
// dp数组大小：[0, n] => n+1
// 返回值：dp[n]

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i-j]*j考虑的是把数字i拆分为多个数字的情况
			// j*(i-j)考虑的是把数字拆分为两个数字的情况
			dp[i] = max(dp[i], dp[i-j]*j, j*(i-j))
		}
	}

	return dp[n]
}

func TestIntegerBreak(t *testing.T) {
	var testData = []struct {
		n    int
		want int
	}{
		{n: 2, want: 1},
		{n: 3, want: 2},
		{n: 10, want: 36},
	}

	for _, tt := range testData {
		get := integerBreak(tt.n)
		if get != tt.want {
			t.Fatalf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
