package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-repeating-substring/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=EASY

// 直接求解
func maxRepeating01(sequence string, word string) int {
	res, idx := 0, 0
	for idx < len(sequence) {
		if idx+len(word) > len(sequence) {
			break
		}
		curr := sequence[idx : idx+len(word)]
		if curr == word {
			res++
			idx += len(word)
			continue
		}
		idx++
	}

	return res
}

func TestMaxRepeating(t *testing.T) {
	var testdata = []struct {
		sequence string
		word     string
		want     int
	}{
		{sequence: "ababc", word: "ab", want: 2},
		{sequence: "ababc", word: "ac", want: 0},
		{sequence: "aaabaaaabaaabaaaabaaaabaaaabaaaaba", word: "aaaba", want: 6},
	}
	for _, tt := range testdata {
		get := maxRepeating01(tt.sequence, tt.word)
		if get != tt.want {
			t.Fatalf("sequence:%v word:%v, want:%v, get:%v", tt.sequence, tt.word, tt.want, get)
		}
	}
}
