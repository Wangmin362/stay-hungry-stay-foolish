package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/repeated-substring-pattern/description/

// 两层for循环,第一层为重复字符串的数量，第二层比较是否是重复的
func repeatedSubstringPattern02(s string) bool {
	for k := 1; k <= len(s)>>1; k++ {
		pattern := s[0:k]
		isValid := true
		for idx := k; idx < len(s); idx += k {
			if idx+k > len(s) || s[idx:idx+k] != pattern { // 说明这个pattern并不是重复的
				isValid = false
				break
			}
		}
		if isValid {
			return true
		}
	}

	return false
}

func TestRepeatedSubstringPattern(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect bool
	}{
		{s: "abab", expect: true},
		{s: "aba", expect: false},
		{s: "aaa", expect: true},
		{s: "abcabc", expect: true},
		{s: "aabaaba", expect: false},
	}

	for _, test := range teatdata {
		get := repeatedSubstringPattern02(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
