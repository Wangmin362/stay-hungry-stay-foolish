package _0_basic

import "testing"

// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 题目分析：找到最长回文子串，所以肯定是连续的
// 明确定义：dp[i][j]表示s[i:j]是否时回文子串
// 递推公式：若s[i]=s[j]  若j-i <=1 dp[i][j] = true  若 j-i>1 dp[i][j] = dp[i+1][j-1]
// 初始化：默认初始化为false即可
// 遍历顺序：从下往上，从左右有
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var maxLen int
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				switch j - i {
				case 0, 1:
					dp[i][j] = true
					if j-i+1 > maxLen {
						maxLen = j - i + 1
						res = s[i : j+1]
					}
				default:
					if dp[i+1][j-1] {
						dp[i][j] = dp[i+1][j-1]
						if j-i+1 > maxLen {
							maxLen = j - i + 1
							res = s[i : j+1]
						}
					}
				}
			}
		}
	}

	return res
}

func TestLongestPalindrome(t *testing.T) {
	var testdata = []struct {
		s    string
		want string
	}{
		{s: "babad", want: "aba"},
		{s: "cbbd", want: "bb"},
	}
	for _, tt := range testdata {
		get := longestPalindrome(tt.s)
		if get != tt.want {
			t.Fatalf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
