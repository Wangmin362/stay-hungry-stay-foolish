package _0_basic

import (
	"testing"
)

// 递归
func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 明确定义：dp[i]就是F(i)的斐波那契数字
// 确定状态转移公式：dp[i] = dp[i-1] + dp[i-2],
// 初始化dp[0]=0, dp[1]=1
// 计算顺序：由于需要当前的数字是从前面的数字推导过来的，因此需要从前到后计算

func fib001(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1) // 需要计算0...n之间的斐波那契数字，因此需要n+1个空间
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 由于dp[i] = dp[i-1] + dp[i-2], 其实i-3之前的缓存完全没有必要，因此这些空间可以不需要记录，我们只需要记录i-1,i-2的数值即可，这样
// 空间复杂度可以优化到O(1), 时间复杂度为O(n)
func fib02(n int) int {
	if n == 0 {
		return 0
	}

	n1, n2 := 0, 1
	// n1, n2, dp
	//     n1, n2, dp
	for i := 2; i <= n; i++ {
		dpi := n1 + n2
		n1 = n2
		n2 = dpi
	}

	return n2
}

// 采用dfs解决：dfs(i) = dfs(i-1) + dfs(i-2)
func fibDfs(n int) int {
	var dfs func(i int) int

	dfs = func(i int) int {
		if i < 2 {
			return i
		}
		return dfs(i-1) + dfs(i-2)
	}
	return dfs(n)
}

// 记忆化搜索
func fibDfsMem(n int) int {
	var dfs func(i int) int

	mem := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mem[i] = -1
	}

	dfs = func(i int) int {
		if i < 2 {
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

// 改为动态规划，也就是递归
// 递归公式为：dfs(i) = dfs(i-1) + dfs(i-2)
// 递推公式为：f[i] = f[i-1] + f[i-2]
// 连边同时加2：f[i+2] = f[i+1] + f[i]
func fibDP(n int) int {
	f := make([]int, n+3)
	for i := 0; i <= n; i++ {
		if i < 2 {
			f[i+2] = i
			continue
		}
		f[i+2] = f[i+1] + f[i]
	}
	return f[n+2]
}

// 改为动态规划，也就是递归
// 递归公式为：dfs(i) = dfs(i-1) + dfs(i-2)
// 递推公式为：f[i] = f[i-1] + f[i-2]
// 连边同时加2：f[i+2] = f[i+1] + f[i]
// 优化空间复杂度
func fibDPOpt(n int) int {
	if n < 2 {
		return n
	}
	f1, f0 := 0, 0
	for i := 0; i <= n; i++ {
		if i < 2 {
			f1 = 1
			f0 = 0
			continue
		}
		f := f1 + f0
		f0 = f1
		f1 = f
	}
	return f1
}
func TestFib(t *testing.T) {
	var testData = []struct {
		n      int
		expect int
	}{
		{n: 3, expect: 2},
		{n: 4, expect: 3},
	}

	for _, tt := range testData {
		get := fibDPOpt(tt.n)
		if tt.expect != get {
			t.Errorf("n:%v, expect:%v, get:%v", tt.n, tt.expect, get)
		}
	}

}
