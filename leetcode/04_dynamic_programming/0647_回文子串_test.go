package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/palindromic-substrings/description/

// 明确定义：dp[i][j]为s[i:j]是否为回文子串
// 递推公式：若s[i]==s[j] if i==j dp[i][j]=true, if j-1=1 dp[i][j]=true, if j-i>1  && dp[i+1][j-1] == true  dp[i][j] = true
// 初始化：全部初始化为false
// 遍历顺序：从下往上，从左往右
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var res int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				switch j - i {
				case 0, 1:
					res++
					dp[i][j] = true
				default:
					if dp[i+1][j-1] {
						res++
						dp[i][j] = true
					}
				}
			}
		}
	}

	return res
}

func TestCountSubstrings(t *testing.T) {
	var testdata = []struct {
		s    string
		want int
	}{
		{s: "abc", want: 3},
		{s: "aaa", want: 6},
	}
	for _, tt := range testdata {
		get := countSubstrings(tt.s)
		if get != tt.want {
			t.Fatalf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
