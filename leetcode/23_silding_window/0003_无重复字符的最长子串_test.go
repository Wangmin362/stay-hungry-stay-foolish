package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	cache := make(map[byte]int) //统计每隔字符的频率
	var res int
	for ; right < len(s); right++ {
		cache[s[right]]++
		for cache[s[right]] > 1 { // 说明当前字符重复 移动左边界，知道不重复
			cache[s[left]]--
			left++
		}
		// 到了这里，当前窗口字符串一定没有重复字符串
		res = max(res, right-left+1)
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
		sum01 := lengthOfLongestSubstring(test.s)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, sum01)
		}
	}
}
