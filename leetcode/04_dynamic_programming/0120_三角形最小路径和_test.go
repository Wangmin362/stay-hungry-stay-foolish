package _0_basic

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/triangle/?envType=study-plan-v2&envId=top-interview-150

// 回溯算法解决
func minimumTotalBacktracking(triangle [][]int) int {
	var backtracking func(x, y, sum int)

	res := math.MaxInt
	path := []int{triangle[0][0]}
	backtracking = func(x, y, sum int) {
		if len(path) == len(triangle) {
			res = min(res, sum)
			return
		}

		if x >= len(triangle)-1 {
			return
		}

		for _, yy := range []int{y, y + 1} {
			path = append(path, triangle[x+1][yy])
			backtracking(x+1, yy, sum+triangle[x+1][yy])
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0, triangle[0][0])
	return res
}

// f(i,j) = min(f(i+1,j), f(i+1,j+1)) + triangle[i][j]
func minimumTotalDigui(triangle [][]int) int {
	var traversal func(x, y int) int

	traversal = func(x, y int) int {
		if x >= len(triangle) {
			return 0
		}
		return min(traversal(x+1, y), traversal(x+1, y+1)) + triangle[x][y]
	}

	return traversal(0, 0)
}

// f(i,j) = min(f(i+1,j), f(i+1,j+1)) + triangle[i][j] 增加记忆化
func minimumTotalDiguiMemo(triangle [][]int) int {
	var traversal func(x, y int) int
	memo := make([][]*int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		memo[i] = make([]*int, len(triangle[len(triangle)-1]))
	}

	traversal = func(x, y int) int {
		if x >= len(triangle) {
			return 0
		}
		if memo[x][y] != nil {
			return *memo[x][y]
		}

		sum := min(traversal(x+1, y), traversal(x+1, y+1)) + triangle[x][y]
		memo[x][y] = &sum

		return sum
	}

	return traversal(0, 0)
}

// 动态规划
// 题目分析：f(i, j)定义为(i,j)点到底边的最小路径和，那么f(i,j) = min(f(i+1,j), f(i+1,j+1)) + triangle(i,j)
// 明确定义：dp[i][j]表示(i,j)到底边的最小路径和
// 递推公式：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle(i,j)
// 初始化：不需要额外的初始化
// 遍历顺序：从底边往上编遍历
func minimumTotal(triangle [][]int) int {
	m, n := len(triangle), len(triangle[len(triangle)-1]) // m+1行，n+1列
	dp := make([][]int, m)
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == m-1 {
				dp[i][j] = triangle[i][j] // 最后一行的最小和路径就是他自己
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}

	return dp[0][0]
}

func TestMinminumTotal(t *testing.T) {
	var testdata = []struct {
		triangle [][]int
		want     int
	}{
		{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, want: 11},
		{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}, want: -1},
	}

	for _, tt := range testdata {
		get := minimumTotal(tt.triangle)
		if get != tt.want {
			t.Fatalf("trangile:%v, want:%v, get:%v", tt.triangle, tt.want, get)
		}
	}
}
