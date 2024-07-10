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
	// dp[n] = dp[n-2] + dp[n-1]
	dp := make([]int, n+1) // 0..n
	dp[0], dp[1], dp[2] = 0, 1, 2
	fmt.Println(dp)
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func climbStairs02(n int) int {
	if n <= 1 {
		return n
	}
	// dp[n] = dp[n-2] + dp[n-1]
	dp := make([]int, n+1) // 0..n
	dp[0], dp[1], dp[2] = 0, 1, 2
	fmt.Println(dp)
	for i := 3; i <= n; i++ {
		for j := 1; j <= 2; j++ { //要么走一步，要么走两步
			dp[i] += dp[i-j]
		}
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
	fmt.Println(climbStairs02(4))
}
