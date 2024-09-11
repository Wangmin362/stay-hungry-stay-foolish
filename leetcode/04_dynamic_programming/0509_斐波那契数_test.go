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

func TestFib(t *testing.T) {
	var testData = []struct {
		n      int
		expect int
	}{
		{n: 3, expect: 2},
		{n: 4, expect: 3},
	}

	for _, tt := range testData {
		get := fib02(tt.n)
		if tt.expect != get {
			t.Errorf("n:%v, expect:%v, get:%v", tt.n, tt.expect, get)
		}
	}

}
