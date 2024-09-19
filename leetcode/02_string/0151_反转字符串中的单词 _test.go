package _1_array

import (
	"slices"
	"strings"
	"testing"
)

// 题目：https://leetcode.cn/problems/reverse-words-in-a-string/

// 先切分，然后反转，最后拼接
func reverseWords01(s string) string {
	split := strings.Split(strings.Trim(s, " "), " ")
	raw := make([]string, 0, len(split))
	for _, str := range split {
		if str != "" {
			raw = append(raw, str)
		}
	}
	slices.Reverse(raw)
	return strings.Join(raw, " ")
}

// 倒序遍历字符串
func reverseWords02(s string) string {
	var res []string
	begin := -1
	for idx := len(s) - 1; idx >= 0; idx-- {
		if s[idx] == ' ' && begin == -1 { // 说明单词还没有开始
			continue
		} else if s[idx] == ' ' && begin != -1 {
			res = append(res, s[idx+1:begin+1])
			begin = -1
		} else if s[idx] != ' ' && begin == -1 {
			begin = idx // 记录单词开始的位置
		}
	}
	if begin != -1 {
		res = append(res, s[0:begin+1])
	}

	return strings.Join(res, " ")
}

// 倒序处理，手动处理两边空格
func reverseWords0919(s string) string {
	var res []string
	idx := len(s) - 1
	first := len(s)
	for idx >= 0 {
		for idx >= 0 && s[idx] == ' ' { // 找到第一个非空字符
			idx--
		}
		if idx < 0 {
			if first == -1 {
				break
			}
			res = append(res, s[idx+1:idx])
			break
		}
		first = idx + 1
		for idx >= 0 && s[idx] != ' ' { // 找到第一个空字符
			idx--
		}
		res = append(res, s[idx+1:first])
		first = -1
	}
	return strings.Join(res, " ")
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
		get := reverseWords0919(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
