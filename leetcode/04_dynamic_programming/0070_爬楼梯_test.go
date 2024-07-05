package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/climbing-stairs/description/

// n = 1 res = 1
// n = 2 res = 2
// n = 3 res = 3
// dp[n] = dp[n-2] + dp[n-1]
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1) // 因为需要保存dp[n]个数字，因此需要n+1个空间
	dp[1], dp[2] = 1, 2
	idx := 3
	for idx <= n {
		dp[idx] = dp[idx-1] + dp[idx-2]
		idx++
	}

	return dp[n]
}

func climbStairs01(n int) int {
	if n <= 1 {
		return n
	}
	dp1, dp2 := 1, 2
	idx := 3
	for idx <= n {
		dpn := dp1 + dp2
		dp1 = dp2
		dp2 = dpn
		idx++
	}

	return dp2
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(5))
}
