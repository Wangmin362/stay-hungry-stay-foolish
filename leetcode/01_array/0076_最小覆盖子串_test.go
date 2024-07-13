package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/minimum-window-substring/description/

func minWindow(s string, t string) string {
	cntStr := func(s string, begin, end int) map[uint8]int {
		m := map[uint8]int{}
		for begin <= end {
			m[s[begin]]++
		}
		return m
	}
	tCnt := cntStr(t, 0, len(t)-1)
	cntCmp := func(c map[uint8]int) bool {
		for k, v := range tCnt {
			cnt, ok := c[k]
			if !ok || cnt < v {
				return false
			}
		}
		return true
	}

	slow, fast := 0, 0
	res := ""
	minLen := len(s) + 1
	for slow < len(s) && fast < len(s) {
		if cntCmp(cntStr(s, slow, fast)) { // 缩小窗口，移动左边界
			length := fast - slow + 1
			if minLen > length {
				minLen = length
				res = s[slow : fast+1]
			}
			slow++
		} else { // 扩大窗口，移动有边界
			fast++
		}
	}

	return res
}

func TestMinWindow(t *testing.T) {
	var testdata = []struct {
		s      string
		t      string
		expect string
	}{
		{s: "ADOBECODEBANC", t: "ABC", expect: "BANC"},
	}

	for _, test := range testdata {
		get := minWindow(test.s, test.t)
		if get != test.expect {
			t.Errorf("s:%v, t:%v  expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}
}
