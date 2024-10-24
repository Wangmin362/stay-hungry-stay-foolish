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

// 递归：dfs(i) = max(dfs(i-1)*i, (i-1)*i, dfs(i-2)*2, (i-2)*2, dfs(i-3)*3, (i-3)*3)
func integerBreakDfs(n int) int {
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = -1
	}

	dfs = func(i int) int {
		if i == 2 { // 比2小的没有意义
			return 1
		}
		if mem[i] != -1 {
			return mem[i]
		}

		res := 0
		for j := 1; j < i; j++ {
			res = max(res, dfs(i-j)*j)
			res = max(res, (i-j)*j)
		}
		mem[i] = res
		return res
	}

	return dfs(n)
}

// 递归：dfs(i) = max(dfs(i-1)*i, (i-1)*i, dfs(i-2)*2, (i-2)*2, dfs(i-3)*3, (i-3)*3)
// 递推：f[i] = max(f[i-1]*i, (i-1)*i, f[i-2]*2, (i-2)*2, f[i-3]*3, (i-3)*3)
func integerBreakDp(n int) int {
	f := make([]int, n+1)
	f[1] = 1
	// f[3] = max(f[2]*1, 2*1, f[1]*2, 1*2
	for i := 2; i <= n; i++ {
		res := 0
		for j := 1; j < i; j++ {
			res = max(res, f[i-j]*j, (i-j)*j)
		}
		f[i] = res
	}
	return f[n]
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
		get := integerBreakDp(tt.n)
		if get != tt.want {
			t.Fatalf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
