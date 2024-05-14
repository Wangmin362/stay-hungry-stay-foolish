package _1_array

import (
	"testing"
)

// 题解思路：https://leetcode.cn/problems/backspace-string-compare/solutions/683776/shuang-zhi-zhen-bi-jiao-han-tui-ge-de-zi-8fn8/

func backspaceCompare(s string, t string) bool {
	sPointer := len(s) - 1
	sSkip := 0 // 用于记录当前需要跳过多少个字符，来一个#就增加一，每遇到一个非#，就减一，相当于抵消一个字符

	tPointer := len(t) - 1
	tSkip := 0
	for sPointer >= 0 || tPointer >= 0 {
		// 找到s字符串第一个需要比较的字符
		for sPointer >= 0 {
			if s[sPointer] == '#' {
				sPointer--
				sSkip++
			} else {
				if sSkip > 0 { // 抵消一个字符
					sPointer--
					sSkip--
				} else {
					// 当前字符不是#号，并且也不需要跳过，那么需要和t字符对比
					break
				}
			}
		}
		for tPointer >= 0 {
			if t[tPointer] == '#' {
				tPointer--
				tSkip++
			} else {
				if tSkip > 0 { // 抵消一个字符
					tPointer--
					tSkip--
				} else {
					// 当前字符不是#号，并且也不需要跳过，那么需要和t字符对比
					break
				}
			}
		}

		if sPointer >= 0 && tPointer >= 0 {
			if s[sPointer] != t[tPointer] {
				return false
			}
		} else if sPointer >= 0 || tPointer >= 0 {
			return false
		}

		tPointer--
		sPointer--
	}

	return true
}

func TestBackspaceCompare(t *testing.T) {
	var testDatas = []struct {
		s1     string
		s2     string
		expect bool
	}{
		{s1: "nzp#o#g", s2: "b#nzp#o#g", expect: true},
		{s1: "ab#c", s2: "ad#c", expect: true},
		{s1: "#####", s2: "###b##", expect: true},
		{s1: "##b", s2: "########b", expect: true},
		{s1: "########b", s2: "##b", expect: true},
	}

	for _, test := range testDatas {
		get := backspaceCompare(test.s1, test.s2)
		if get != test.expect {
			t.Errorf("s1:%s, s2:%s, expect:%v, get:%v", test.s1, test.s2, test.expect, get)
		}
	}
}
