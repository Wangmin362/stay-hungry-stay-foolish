package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLongestSubstring(s string) int {
	m := map[rune]struct{}{}
	c := list.New()
	maxLength := 0
	for _, char := range s {
		_, ok := m[char]
		if ok { // 重复出现，移除掉重复元素之前的元素
			for {
				cc := c.Remove(c.Front()).(rune)
				if cc == char { // 找到了那个需要移除的元素
					c.PushBack(char)
					break
				} else {
					delete(m, cc)
				}
			}

		} else {
			c.PushBack(char)
			if c.Len() > maxLength {
				maxLength = c.Len()
			}
			m[char] = struct{}{}
		}
	}

	return maxLength
}

func lengthOfLongestSubstring01(s string) int {
	m := make(map[rune]int) // 记录当前字符，以及当前字符最近一次出现的索引位置
	left, maxLen := 0, 0
	for i, char := range s {
		if idx, ok := m[char]; ok { // 一旦当前字符以前出现过，那么窗口的左边界只能移动到这个重复字符的位置
			if idx >= left { // 只有重复元素在窗口内部才需要移动左边界
				left = idx + 1
			}
		}
		m[char] = i // 记录元素是否出现过
		maxLen = max(maxLen, i-left+1)
	}
	return maxLen
}

func TestThreeSum(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect int
	}{
		{s: "abcabcabc", expect: 3},
		{s: "abba", expect: 2},
		{s: "pwwkew", expect: 3},
		{s: "dvdf", expect: 3},
	}

	for _, test := range teatdata {
		sum01 := lengthOfLongestSubstring01(test.s)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, sum01)
		}
	}
}
