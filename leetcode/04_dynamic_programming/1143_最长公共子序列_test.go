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

/*
///////////////// dfs(i-1, j-1)+1 s[i] = s[j]
递归：dfs(i,j) =
///////////////// max(dfs(i, j-1), dfs(i-1,j))  s[i]!=s[j]
*/
func longestCommonSubsequenceDfs(text1 string, text2 string) int {
	var dfs func(i, j int) int

	mem := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		mem[i] = make([]int, len(text2))
		for j := 0; j < len(text2); j++ {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}

		if text1[i] == text2[j] {
			res := dfs(i-1, j-1) + 1
			mem[i][j] = res
			return res
		} else {
			res := max(dfs(i, j-1), dfs(i-1, j))
			mem[i][j] = res
			return res
		}
	}
	return dfs(len(text1)-1, len(text2)-1)
}

/*
///////////////// dfs(i-1, j-1)+1 s[i] = s[j]
递归：dfs(i,j) =
///////////////// max(dfs(i, j-1), dfs(i-1,j))  s[i]!=s[j]
改为递推：
///////////////// f[i-1][j-1]+1 s[i] = s[j]
递归：f[i][j] =
///////////////// max(f[i][j-1], f[i-1][j])  s[i]!=s[j]
// 为了防止负数下标，两边同时加一，可得：
///////////////// f[i][j]+1 s[i] = s[j]
递归：f[i+1][j+1] =
///////////////// max(f[i+1][j], f[i][j+1])  s[i]!=s[j]
*/
func longestCommonSubsequenceDp(text1 string, text2 string) int {
	f := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			f[i] = make([]int, len(text2)+1)
		}
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}

	return f[len(text1)][len(text2)]
}

/*
///////////////// dfs(i-1, j-1)+1 s[i] = s[j]
递归：dfs(i,j) =
///////////////// max(dfs(i, j-1), dfs(i-1,j))  s[i]!=s[j]
改为递推：
///////////////// f[i-1][j-1]+1 s[i] = s[j]
递归：f[i][j] =
///////////////// max(f[i][j-1], f[i-1][j])  s[i]!=s[j]
// 为了防止负数下标，两边同时加一，可得：
///////////////// f[i][j]+1 s[i] = s[j]
递归：f[i+1][j+1] =
///////////////// max(f[i+1][j], f[i][j+1])  s[i]!=s[j]

// 优化为两行可得：
///////////////// f[i%2][j]+1 s[i] = s[j]
递归：f[(i+1)%2][j+1] =
///////////////// max(f[(i+1)%2][j], f[i%2][j+1])  s[i]!=s[j]
*/
func longestCommonSubsequenceDpOpt1(text1 string, text2 string) int {
	f := make([][]int, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j <= len(text2); j++ {
			f[i] = make([]int, len(text2)+1)
		}
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				f[(i+1)%2][j+1] = f[i%2][j] + 1
			} else {
				f[(i+1)%2][j+1] = max(f[(i+1)%2][j], f[i%2][j+1])
			}
		}
	}

	return f[len(text1)%2][len(text2)]
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
		get := longestCommonSubsequenceDpOpt1(tt.text1, tt.text2)
		if get != tt.want {
			t.Fatalf("text1:%v, text2:%v, want:%v, get:%v", tt.text1, tt.text2, tt.want, get)
		}
	}
}
