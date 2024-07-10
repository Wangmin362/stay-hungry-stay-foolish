package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/unique-paths/description/

func uniquePaths(m int, n int) int {
	// dp[i][j]为机器人走到(i,j)的路径
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, m) // [0,m]
	dp[0] = make([]int, n)
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	fmt.Println(dp[0])
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
		fmt.Println(dp[i])
	}

	return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(3, 2))
}
