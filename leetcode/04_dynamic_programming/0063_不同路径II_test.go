package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/unique-paths-ii/description/

// 明确定义：dp[i][j]表示机器人移动到(i,j)坐标的不同路径
// 转移方程：dp[i][j] = dp[i][j-1] + dp[i-1][j]
// 初始化： dp[0][j] = 1, dp[i][0] = 1 需要排除障碍物
// 遍历顺序：从上往下，从左往右
// dp数组大小： dp[m][n]
// 返回值：dp[m-1][n-1]

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	// 第一列初始化
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 { // 遇到障碍物直接，返回，后续不可能走到
			break
		}
		dp[i][0] = 1
	}

	// 第一行初始化
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var testData = []struct {
		obs  [][]int
		want int
	}{
		{obs: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: 2},
		{obs: [][]int{{0, 1}, {0, 0}}, want: 1},
	}

	for _, tt := range testData {
		get := uniquePathsWithObstacles(tt.obs)
		if get != tt.want {
			t.Fatalf("obs:%v, want:%v, get:%v", tt.obs, tt.want, get)
		}
	}
}
