package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/perfect-squares/description/

func numSquares(n int) int {
	// dp[j]定义为1...i个数字装满背包容量为j的最小数量
	// dp[j] = min(dp[j], dp[j - nums[i]]+1)
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt
	}

	for j := 1; j <= n; j++ {
		for i := 1; i*i <= j; i++ {
			if dp[j-i*i] != math.MaxInt {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-i*i]+1)))
			}
		}
		fmt.Println(dp)
	}
	return dp[n]
}

func numSquares01(n int) int {
	// dp[j]定义为1...i个数字装满背包容量为j的最小数量
	// dp[j] = min(dp[j], dp[j - nums[i]]+1)
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j-i*i] != math.MaxInt {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-i*i]+1)))
			}
		}
		fmt.Println(dp)
	}
	return dp[n]
}

/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
/*45*/
func numSquares03(n int) int {
	// dp[j] = min(dp[j], dp[j-i*i] + 1)
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-i*i]+1)))
		}
		fmt.Println(dp)
	}

	return dp[n]
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares03(12))
}
