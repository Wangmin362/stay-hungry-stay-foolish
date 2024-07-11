package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/is-subsequence/description/

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 || len(s) > len(t) {
		return false
	}

	// dp[i][j]定义为s[:i-1]和t[:j-1]的最长公共子序列
	dp := make([][]int, len(s)+1)
	dp[0] = make([]int, len(t)+1)
	res := 0
	for i := 1; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res == len(s)
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
