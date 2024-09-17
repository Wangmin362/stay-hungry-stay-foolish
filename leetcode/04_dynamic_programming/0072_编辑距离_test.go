package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/edit-distance/description/

// 题目分析：求word1删除、新增、替换字符之后变为word2的最小次数
// 明确定义：dp[i][j]表示以i-1结尾的word1和以j-1为结尾的word2的最小删除次数
// 递推公式：若word1[i-1]==word2[j-1]，那么dp[i][j] = dp[i-1][j-1]，因为已经相等，不需要操作
// 若word1[i-1]!=word[j-1]，那么可以删除word1,即dp[i-1][j]+1, 可以删除word2,即dp[i][j-1]+1, 可以两个同时删除即dp[i-1][j-1]+2
// 可以可以替换即dp[i-1][j-1]+1，综上dp[i][j]=min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2, dp[i-1][j-1]+1)
// 由于dp[i-1][j-1]+2 一定大于 dp[i-1][j-1]+1，因此 dp[i][j]=min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
// 初始化：dp[i][0] = i, dp[0][j] = j
func minDistanceII(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func TestMinDistanceII(t *testing.T) {
	var testdata = []struct {
		word1 string
		word2 string
		want  int
	}{
		{word1: "horse", word2: "ros", want: 3},
		{word1: "intention", word2: "execution", want: 5},
	}

	for _, tt := range testdata {
		get := minDistanceII(tt.word1, tt.word2)
		if get != tt.want {
			t.Fatalf("word1:%v, word2:%v, want:%v, get:%v", tt.word1, tt.word2, tt.want, get)
		}
	}
}
