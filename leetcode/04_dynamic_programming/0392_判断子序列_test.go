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

// 明确定义：dp[i][j]表示以i-1几位的s字符串和以j-1结尾的最长公共子序列的最大长度
// 递推公式：若s[i-1] = t[j-1]，那么dp[i][j] = dp[i-1][j-1]+1
// 若s[i-1] != t[j-1]，那么dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // TODO 实际上，这里不应该删除s，之应该删除t，因为判断的是s是不是t的子序列
// TODO 所以 若s[i-1] != t[j-1]，那么dp[i][j] =  dp[i][j-1] ，这里之应该删除t
// 初始化：dp[0][j] = 0, dp[i][j] = 0
// 遍历顺序：从小到大，从左往右
func isSubsequence02(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// TODO 实际上，这里不应该删除s，之应该删除t，因为判断的是s是不是t的子序列
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

// 使用双指针，一个指针指向s，一个指针指向t，只有当两个指针指向的字符相同，才同时移动指针，否则移动指向t的指针，因为需要在t中找到一个和s相同的字符
func isSubsequence03(s string, t string) bool {
	sidx, tidx := 0, 0
	for sidx < len(s) && tidx < len(t) {
		if s[sidx] == t[tidx] {
			sidx++
		}

		tidx++
	}

	return sidx == len(s)
}

func TestIsSubsequence(t *testing.T) {
	var testdata = []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abc", t: "ahbgdc", want: true},
		{s: "axc", t: "ahbgdc", want: false},
	}

	for _, tt := range testdata {
		get := isSubsequence03(tt.s, tt.t)
		if get != tt.want {
			t.Fatalf("s:%v, t:%v, want:%v, get:%v", tt.s, tt.t, tt.want, get)
		}
	}

}
