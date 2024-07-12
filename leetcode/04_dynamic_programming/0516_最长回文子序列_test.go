package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-palindromic-subsequence/description/

func longestPalindromeSubseq(s string) int {
	// dp[i][j]为s[i:j]包含j的最长回文子序列的长度
	// s[i] == s[j]  dp[i][j] = dp[i+1][j-1]+2
	// s[i] != s[j] dp[i][j] = max(dp[i+1][j], dp[i][j-1])
	dp := make([][]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(s))
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if i == j {
					dp[i][j] = 1
				} else if j-i == 1 {
					dp[i][j] = 2
				} else if j-i > 1 {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
		fmt.Println(dp[i])
	}

	return dp[0][len(s)-1]
}

func TestLongestPalindromeSubseq(t *testing.T) {
	fmt.Println(longestPalindromeSubseq("aaa"))
}
