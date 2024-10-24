package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/unique-binary-search-trees/description/

// 题目分析：n=3,所有BST的数量为根节点为1的BST数量 + 根节点为2的BST数量 + 根节点为3的BST数量
// 根节点为1的BST数量= dp[0] * dp[2]
// 根节点为2的BST数量= dp[1] * dp[1]
// 根节点为3的BST数量= dp[2] * dp[0]
// 明确定义：dp[i]表示i个节点组成的BST树的数量
// 转移方程：dp[i] = dp[0]*dp[i-1] + dp[1]*dp[i-2] + dp[2]*dp[i-3] + ... +  dp[i-2]*dp[1] + dp[i-1]*dp[1]
// 初始化： dp[0] = 1, dp[1] = 1, dp[2] = 2
// 遍历顺序：从前往后
// dp数组大小：[0, n] => n+1
// 返回值：dp[n]

func numTrees(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

// 递归：dfs(n) = dfs(0)*dfs(n-1) + dfs(1)*dfs(n-2) + dfs(2)*dfs(n-3) + dfs(3)*dfs(n-4)
// dfs(0)=1, dfs(1)=1
func numTreesDfs(n int) int {
	var dfs func(i int) int
	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		if mem[i] != -1 {
			return mem[i]
		}

		var res int
		for j := 0; j < i; j++ {
			res += dfs(j) * dfs(i-j-1)
		}
		mem[i] = res
		return res
	}
	return dfs(n)
}

func TestNumTrees(t *testing.T) {
	var testData = []struct {
		n    int
		want int
	}{
		{n: 3, want: 5},
		{n: 1, want: 1},
	}

	for _, tt := range testData {
		get := numTreesDfs(tt.n)
		if get != tt.want {
			t.Fatalf("n:%v, want:%v, get:%v", tt.n, tt.want, get)
		}
	}
}
