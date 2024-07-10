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

// 排列问题  完全背包  TODO
func wordBreak02(s string, wordDict []string) bool {
	// dp[j]定义为前j个字符是否可以有wordDict凭借而成
	// dp[j] = dp[j-i] && s[i:j] in wordDict
	wMap := map[string]struct{}{}
	for idx := range wordDict {
		wMap[wordDict[idx]] = struct{}{}
	}
	dp := make([]bool, len(s)+1) // 默认初始化为false
	dp[0] = true
	for j := 0; j <= len(s); j++ {
		for i := 0; i < j; i++ {
			str := wordDict[i]
			if j >= len(str) {
				ss := s[j-len(str) : j]
				_, ok := wMap[ss]
				dp[j] = dp[j-len(str)] && ok
			}
		}
		fmt.Println(dp)
	}

	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	//fmt.Println(wordBreak02("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak02("applepenapple", []string{"apple", "pen"}))
}
