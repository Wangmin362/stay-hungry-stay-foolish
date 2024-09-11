package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/unique-paths/description/

// 题目分析：由于机器人只能向前向下移动，因此要想达到某个点，只能从上方或者前方移动过来，对应第一行和第一列，肯定是0，因为只能向前移动或者向后移动
// 明确定义：dp[i][j]表示机器人移动到(i,j)坐标的不同路径
// 转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 初始化：dp[i][0] = 1  dp[0][j] = 1  dp[0][0]没有意义
// 遍历方向：从前往后，从上往下
// dp数组大小 dp[m][n]
// 返回值：dp[m-1][n-1]

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // 第一列初始化为1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 // 第一行初始化为1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

// 状态压缩，每个位置只需要知道当前行与上一行的状态即可，也就是说只需要两行数组即可搞定
func uniquePaths01(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		dp[1][0] = 1 // 第一列等于1
		for j := 1; j < n; j++ {
			dp[1][j] = dp[1][j-1] + dp[0][j]
		}
		dp[0] = dp[1]
	}

	return dp[1][n-1]
}

func TestUniquePaths(t *testing.T) {
	var testData = []struct {
		m    int
		n    int
		want int
	}{
		{m: 1, n: 1, want: 1},
		{m: 2, n: 2, want: 2},
		{m: 3, n: 7, want: 28},
		{m: 3, n: 2, want: 3},
		{m: 3, n: 3, want: 6},
	}

	for _, tt := range testData {
		get := uniquePaths01(tt.m, tt.n)
		for get != tt.want {
			t.Fatalf("m:%v, n:%v, want:%v, get:%v", tt.m, tt.n, tt.want, get)
		}
	}
}
