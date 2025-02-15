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

// 递归：dfs(i) = dfs(i-1) + dfs(i-2)
func climbStairsDfs(n int) int {
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i <= 2 {
			return i
		}
		if mem[i] != -1 {
			return mem[i]
		}
		res := dfs(i-1) + dfs(i-2)
		mem[i] = res
		return res
	}
	return dfs(n)
}

// 递归：dfs(i) = dfs(i-1) + dfs(i-2)
// 递推：f[i] = f[i-1] + f[i-2]
func climbStairsDP(n int) int {
	f := make([]int, n+1)
	f[0], f[1], f[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 递归：dfs(i) = dfs(i-1) + dfs(i-2)
// 递推：f[i] = f[i-1] + f[i-2]
// 空间优化
func climbStairsDPOpt(n int) int {
	if n <= 2 {
		return n
	}

	f1, f0 := 2, 1
	for i := 3; i <= n; i++ {
		f := f1 + f0
		f0 = f1
		f1 = f
	}
	return f1
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
		get := climbStairsDPOpt(tt.n)
		if get != tt.want {
			t.Errorf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
