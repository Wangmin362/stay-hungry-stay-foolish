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

func lengthOfLongestSubstring0902(s string) int {
	cache := make(map[byte]int, len(s))
	ans, left := 0, 0
	for right := 0; right < len(s); right++ {
		cache[s[right]]++
		if right-left+1 == len(cache) {
			ans = max(ans, len(cache))
		} else {
			for right-left+1 > len(cache) { // 说明有重复元素，此时需要移动左端点
				if cnt := cache[s[left]]; cnt == 1 {
					delete(cache, s[left])
				} else {
					cache[s[left]]--
				}
				left++
			}
		}
	}
	return ans
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
		sum01 := lengthOfLongestSubstring0902(test.s)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, sum01)
		}
	}
}
