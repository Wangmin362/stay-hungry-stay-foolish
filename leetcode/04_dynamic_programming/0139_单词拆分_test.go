package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/word-break/description/

func wordBreak(s string, wordDict []string) bool {
	// dp[j]定义为容量为j长度的背包，是否可以把wordDict装满，由于是把数组中拼接为字符串，因此是有顺序的，因此是一个排列问题
	// dp[j] = str[i:j] == wordDict[i] && dp[j-wordDict[i]]
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for j := 1; j <= len(s); j++ { // 遍历背包
		for i := 0; i < j; i++ { // 遍历物品
			if _, ok := wordMap[s[i:j]]; ok && dp[i] {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
