package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/word-break/description/
func wordBreakBacktracking(s string, wordDict []string) bool {
	m := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		m[word] = struct{}{}
	}

	var res bool
	var backtracking func(start int)
	backtracking = func(start int) {
		if res {
			return
		}
		if start >= len(s) {
			res = true
			return
		}

		for i := start; i < len(s); i++ {
			if _, ok := m[s[start:i+1]]; !ok {
				continue
			}
			backtracking(i + 1)
		}
	}

	backtracking(0)
	return res
}

// 题目分析：背包的容量为s, 物品为wordDict, 每个物品可以取用无数次，那么是一个完全背包问题。因此背包从小到大遍历，由于题目没有强调顺序
// 因此先背包，再物品，或者先物品，在背包都可以
// 明确定义：dp[j]表示s[0,j+1]是否可以由单词拼接出来
// 递推公式：如果已知dp[j]可以拼接出来，并且s[j:i]再单词表中，那么dp[i]也可以拼接出来
// 初始化dp[0] = true
// 遍历顺序：本体虽然，没有强调顺序，但是由于是拼接字符串，所以是存在先后顺序的，因此先背包，后物品
func wordBreak(s string, wordDict []string) bool {
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
	var testdata = []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{s: "leetcode", wordDict: []string{"leet", "code"}, want: true},
		{s: "applepenapple", wordDict: []string{"apple", "pen"}, want: true},
	}
	for _, tt := range testdata {
		get := wordBreak(tt.s, tt.wordDict)
		if get != tt.want {
			t.Fatalf("s:%v, wordDict:%v, want:%v, get:%v", tt.s, tt.wordDict, tt.want, get)
		}
	}
}
