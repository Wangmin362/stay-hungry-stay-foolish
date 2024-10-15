package main

import (
	"testing"
)

// 状态定义：dp[i][j]表示s[i:j]回文子串的数量
// 递推公式：
// 若s[i] == s[j]，那么dp[i][j] = dp[i+1][j-1] + 2
// 若s[i] != s[j]，那么dp[i][j] = max(dp[i][j-1], dp[i-1][j])
// 初始化：dp[i][i]=1
func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	var res int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}

func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
}
