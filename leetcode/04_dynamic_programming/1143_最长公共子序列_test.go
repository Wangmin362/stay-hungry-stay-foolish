package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/longest-common-subsequence/description/

// 题目分析：求出text1和text2字符串的最长公共子序列，子序列可以是不连续的
// 明确定义：dp[i][j]表示以text1[0:i-1]和text2[0:j-1]的最长公共子序列的长度
// 递推公式：若text1[i-1]==text2[j-1] 那么 dp[i][j] = dp[i-1][j-1]+1, 否则dp[i][j] = max(dp[i][j-1], dp[i-1][j])的子序列
// 初始化：dp[i][0] = 0, dp[0][j] = 0
// dp数组大小，0没有意义，数组一直到len(text1)，那就是len(text1)+1
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	var res int
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}

func TestLongestCommonSubsequence(t *testing.T) {
	var testdata = []struct {
		text1 string
		text2 string
		want  int
	}{
		{text1: "abcde", text2: "ace", want: 3},
	}
	for _, tt := range testdata {
		get := longestCommonSubsequence(tt.text1, tt.text2)
		if get != tt.want {
			t.Fatalf("text1:%v, text2:%v, want:%v, get:%v", tt.text1, tt.text2, tt.want, get)
		}
	}
}
