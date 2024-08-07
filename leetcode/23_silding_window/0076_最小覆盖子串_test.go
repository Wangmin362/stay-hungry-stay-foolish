package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/minimum-window-substring/description/

// 滑动窗口，只不过这种方式时间超长
func minWindow(s string, t string) string {
	cntStr := func(s string, begin, end int) map[uint8]int {
		m := map[uint8]int{}
		for begin <= end {
			m[s[begin]]++
			begin++
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

// 滑动窗口
func minWindow01(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}
	sm, st := [128]int{}, [128]int{}

	isCover := func() bool {
		for char, cnt := range st {
			if sm[char] < cnt {
				return false
			}
		}
		return true
	}

	for _, char := range t {
		st[char]++
	}

	left, right := 0, 0
	minLen := len(s) + 1
	resL, resR := -1, minLen
	for right = range s {
		sm[s[right]]++
		for isCover() {
			if right-left+1 < minLen {
				minLen = right - left + 1
				resL, resR = left, right
			}
			sm[s[left]]--
			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[resL : resR+1]
}

func TestMinWindow(t *testing.T) {
	testdata := []struct {
		s      string
		t      string
		expect string
	}{
		{s: "ADOBECODEBANC", t: "ABC", expect: "BANC"},
		{s: "aa", t: "a", expect: "a"},
		{s: "a", t: "aa", expect: ""},
	}

	for _, test := range testdata {
		get := minWindow01(test.s, test.t)
		if get != test.expect {
			t.Errorf("s:%v, t:%v  expect:%v, get:%v", test.s, test.t, test.expect, get)
		}
	}
}
