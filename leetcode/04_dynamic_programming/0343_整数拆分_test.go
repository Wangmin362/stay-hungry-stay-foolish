package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/integer-break/description/

func integerBreak(n int) int {
	// 假设n = k1 + k2 + k3 + k4
	// 那么n一定可以拆分为1*(n-1) 2*(n-2) 3*(n-3) 4*(n-4)
	// 那么dp[i] = max(x*(i-x)) x=[1,n-1]
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		iMax := 0
		for x := 1; x < i; x++ {
			iMax = max(iMax, x*dp[i-x]) // 拆分N个数字
			iMax = max(iMax, x*(i-x))   // 拆分两个数字
		}
		dp[i] = iMax
		fmt.Println(dp)
	}

	return dp[n]
}

func TestIntegerBreak(t *testing.T) {
	fmt.Println(integerBreak(4))
}
