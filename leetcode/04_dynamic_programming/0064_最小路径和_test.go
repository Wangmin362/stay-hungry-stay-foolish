package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/minimum-path-sum/description/?envType=study-plan-v2&envId=top-interview-150

// 明确定义：dp[i][j]表示之前走过的最小路径和
// 递推公式：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + gird[i][j]
// 初始化：dp[0][j] = gird[0][j]  dp[i][0] = gird[i][0]
// 遍历顺序：从上往下，从左往右
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func TestMinPathSum(t *testing.T) {
	var testdata = []struct {
		grid [][]int
		want int
	}{
		{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, want: 7},
	}

	for _, tt := range testdata {
		get := minPathSum(tt.grid)
		if get != tt.want {
			t.Fatalf("grid:%v, want:%v, get:%v", tt.grid, tt.want, get)
		}
	}
}
