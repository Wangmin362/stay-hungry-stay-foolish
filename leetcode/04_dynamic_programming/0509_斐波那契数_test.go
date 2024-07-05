package _0_basic

import (
	"fmt"
	"testing"
)

// 递归
func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 动态规划
func fib02(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1) // dp[n]表示n的斐波那契数字，因为要存储[0, n]一共n+1个数字的斐波那契数字，因此数字大小为n+1
	// 递推公式 dp[n] = dp[n-2] + dp[n-1]
	dp[0], dp[1] = 0, 1
	idx := 2
	for idx <= n {
		dp[idx] = dp[idx-2] + dp[idx-1]
		idx++
	}
	return dp[n]
}

// 动态规划精简版本  本质上，我们只需要记录dp[n-1]和dp[n-2]的指，不需要记录n-2之前的值，因为没有要求返回这些数字，而且这些数字也没有任何作用
func fib03(n int) int {
	if n == 0 {
		return 0
	}
	dp1, dp2 := 0, 1 // dp1表示dp[n-1], dp2表示dp[n-2]
	// 递推公式 dp[n] = dp[n-2] + dp[n-1]
	idx := 2
	for idx <= n {
		dpn := dp1 + dp2
		dp1 = dp2
		dp2 = dpn
		idx++
	}

	return dp2
}

func TestFib(t *testing.T) {
	fmt.Println(fib02(5))
}
