package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/distinct-subsequences/description/

// 题目分析：找到s中有多少个t，其实就是删除s中的元素，看看是否能够等于t, 相当于求删除s的有多少种方法可以变为元素t
// 明确定义：dp[i][j]表示以i-1为结尾的s有多少个以j-1为结尾的t
// 递推公式：if s[i-1]==t[j-1] dp[i][j] = dp[i-1][j-1] + dp[i-1][j]  else dp[i][j] = dp[i-1][j]
// 初始化：dp[0][j] = 0, dp[i][0] = 1, dp[0][0] = 1
// dp数组大小： [len(s)+1][len(t)+1]
// 返回值：dp[len(s)][len(t)]
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]
}

func TestNumDistinct(t *testing.T) {
	var testdata = []struct {
		s    string
		t    string
		want int
	}{
		{s: "babgbag", t: "bag", want: 5},
		{s: "rabbbit", t: "rabbit", want: 3},
	}
	for _, tt := range testdata {
		get := numDistinct(tt.s, tt.t)
		if get != tt.want {
			t.Fatalf("s:%v, t:%v, want:%v, get:%v", tt.s, tt.t, tt.want, get)
		}
	}
}
