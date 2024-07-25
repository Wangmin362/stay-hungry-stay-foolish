package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

func maxVowels(s string, k int) int {
	isVowel := func(char uint8) bool {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			return true
		}
		return false
	}

	k = min(len(s), k)
	maxLen := 0
	currLen := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currLen++
		}
	}
	maxLen = max(maxLen, currLen)
	left, right := 0, k-1
	for right < len(s) {
		right++
		if right < len(s) && isVowel(s[right]) {
			currLen++
		}
		if isVowel(s[left]) {
			currLen--
		}
		left++
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
func TestMaxVowels(t *testing.T) {
	testdata := []struct {
		s      string
		t      int
		expect int
	}{
		{s: "abciiidef", t: 3, expect: 3},
	}

	for _, test := range testdata {
		get := maxVowels(test.s, test.t)
		if get != test.expect {
			t.Errorf("s:%v, t:%v  expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}
}
