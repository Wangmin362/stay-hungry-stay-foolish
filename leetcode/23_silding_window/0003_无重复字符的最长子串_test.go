package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150

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

func lengthOfLongestSubstring03(s string) int {
	left, right := 0, 0 // 左右边界
	cache := make([]int, 256)
	var res int
	for right < len(s) {
		cache[s[right]]++         // 判断当前字符是否重复，如有重复，移动左边界，直到不重复
		if cache[s[right]] <= 1 { // 说明没有重复，计算最大值
			res = max(res, right-left+1)
		} else {
			for cache[s[right]] > 1 { // 移动左边界，直到没有重复字符串
				cache[s[left]]--
				left++
			}
		}
		right++
	}

	return res
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
		sum01 := lengthOfLongestSubstring03(test.s)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, sum01)
		}
	}
}
