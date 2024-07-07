package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/integer-break/description/

func integerBreak(n int) int {
	// dp[i]为正整数i拆分为k个数字，并且数字乘积最大化结果
	// dp[i] = dp[j] * dp[i-j]
	// dp[0] = 0  dp[1] = 1 dp[2] = 1
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		maxDp := math.MinInt
		for j := 1; j < i; j++ {
			dp2 := j * (i - j) // 拆分为两个数字
			dpn := j * dp[i-j] // 拆分为三个以及三个数字以上
			if dpn > maxDp {
				maxDp = dpn
			}
			if dp2 > maxDp {
				maxDp = dp2
			}
		}
		dp[i] = maxDp
	}

	return dp[n]
}

func TestIntegerBreak(t *testing.T) {
	fmt.Println(integerBreak(10))
}
