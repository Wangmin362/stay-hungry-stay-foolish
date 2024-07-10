package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/unique-paths-ii/description/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) <= 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// dp[i][j] = dp[i][j-1] + dp[i-1][j]
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			break
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}

	fmt.Println(dp[0])
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
		fmt.Println(dp[i])
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0}, {1, 1}, {0, 0}}))
}
