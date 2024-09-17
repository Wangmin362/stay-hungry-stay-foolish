package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/is-subsequence/description/

// 题目分析：求出s和t的最长子序列，判断最长子序列是否是s的长度，如果是，那就是
// 明确定义：dp[i][j]表示s[0:i-1]和t[0:j-1]的最长子序列
// 递推公式：if s[i-1]==t[j-1] dp[i][j]=dp[i-1][j-1]+1  else  dp[i][j] = dp[i][j-1] /
// 初始化：dp[0][j],dp[i][0] = 0,0
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1] // 这里只需要删除t，不需要删除s，因为是要判断s是否是t的子串
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

func TestIsSubsequence(t *testing.T) {
	var testdata = []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abc", t: "ahbgdc", want: true},
		{s: "a3c", t: "ahbgdc", want: false},
	}

	for _, tt := range testdata {
		get := isSubsequence(tt.s, tt.t)
		if get != tt.want {
			t.Fatalf("s:%v, t:%v, want:%v, get:%v", tt.s, tt.t, tt.want, get)
		}
	}

}
