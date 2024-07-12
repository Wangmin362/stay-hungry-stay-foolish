package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/palindromic-substrings/description/

func countSubstrings(s string) int {
	// dp[i][j]定义为s[i:j]包含j的回文子串的数量
	// s[i] == s[j]
	//    若 i=j , dp[i][j] = true
	//    若 j-i=1, dp[i][j] = true
	//    若 j-i>1, dp[i][j] = dp[i+1][j-1]
	// s[i] != s[j]  dp[i][j] = false
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}

	return res
}

func TestCountSubstrings(t *testing.T) {
	fmt.Println(countSubstrings("aaaaa"))
}
