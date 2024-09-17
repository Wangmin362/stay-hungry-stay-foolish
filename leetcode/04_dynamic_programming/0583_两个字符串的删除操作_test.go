package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/delete-operation-for-two-strings/description/

// 题目分析：任意删除word1和word2的字符，是的它们相等的最小操作次数
// 明确定义：dp[i][j]为i-1为结尾的word1和以j-1为结尾的word2相等需要删除元素的最小操作次数
// 地推公式： 当word1[i-1]==word2[j-1]时，不需要操作，此时dp[i][h] = dp[i-1][j-1]
// 当word1[i-1]!=word2[j-1]时，我们可以考虑删除word1的元素, 即dp[i-1][j] + 1, 也可以考虑删除word2字符串，即dp[i][j-1]+1
// 也可以考虑同时删除word1,word2，此时dp[i-1][j-1]+2，综上：dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
// 初始化：根据定义dp[i][0] = i, 因为word1有i个元素，word2有0个元素，要想两个相等，显然只能是word1删除i个元素
// 同理可得：dp[0][j] = j
// dp数组大小: [len(word1)+1][len(word2)]+1
// 返回值：dp[len(word1)][len(word2)]
func minDistance(word1 string, word2 string) int {
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
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func TestMinDistance(t *testing.T) {
	var testdata = []struct {
		word1 string
		word2 string
		want  int
	}{
		{word1: "sea", word2: "eat", want: 2},
	}
	for _, tt := range testdata {
		get := minDistance(tt.word1, tt.word2)
		if get != tt.want {
			t.Fatalf("word1:%v, word2:%v, want:%v, get:%v", tt.word1, tt.word2, tt.want, get)
		}
	}
}
