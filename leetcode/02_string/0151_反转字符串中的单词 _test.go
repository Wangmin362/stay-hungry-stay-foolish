package _1_array

import (
	"strings"
	"testing"
)

// 题目：https://leetcode.cn/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	s = strings.Trim(s, " ") // 先去掉左右两边空格

	res := ""
	endIdx := len(s) - 1
	for idx := len(s) - 1; idx >= 0; idx-- {
		if s[idx] == ' ' { // 第一次遇到空格
			res += s[idx+1 : endIdx+1]      // 取出单词
			for idx >= 0 && s[idx] == ' ' { // 去除多余的空格
				idx--
			}
			endIdx = idx
			if idx >= 0 {
				res += " "
			}
		}
	}
	res += s[:endIdx+1]

	return res
}

func TestReverseWords(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect string
	}{
		{s: " the sky is blue ", expect: "blue is sky the"},
		{s: "      hello    world", expect: "world hello"},
	}

	for _, test := range teatdata {
		get := reverseWords(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
