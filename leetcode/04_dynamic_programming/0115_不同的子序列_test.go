package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/distinct-subsequences/description/

func numDistinct(s string, t string) int {
	// dp[i][j]定义为s[:i-1]字符串在t[:i-1]中出现的次数
	// s[i-1] = t[j-1]时， dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// s[i-1] != t[j-1]时，dp[i][j] = dp[i-1][j]
	dp := make([][]int, len(s)+1)
	dp[0] = make([]int, len(t)+1)
	for j := 0; j <= len(t); j++ {
		dp[0][j] = 0
	}
	dp[0][0] = 1 // 空字符串在空字符串中可以出现一次
	fmt.Println(dp[0])
	for i := 1; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
		fmt.Println(dp[i])
	}

	return dp[len(s)][len(t)]
}

func TestNumDistinct(t *testing.T) {
	fmt.Println(numDistinct("babgbag", "bag"))
}
