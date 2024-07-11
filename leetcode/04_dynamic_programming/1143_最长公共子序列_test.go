package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-common-subsequence/description/

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j]定义为前i-1个text1字符串和前j-1个text2长度的
	// 如果text1[i-1]=text2[j-1] dp[i][j] = dp[i-1][j-1] + 1
	// 如果text[i-1]!=text2[j-1]，那么我们就需要看text[0:i-2]和text1[0:j]以及text[0:i-1]和text2[0:j-2]的最长公共子序列
	// 即dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)
	res := 0
	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
