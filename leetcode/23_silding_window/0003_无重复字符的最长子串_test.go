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
func TestThreeSum(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect int
	}{
		{s: "abcabcabc", expect: 3},
		{s: "abba", expect: 2},
	}

	for _, test := range teatdata {
		sum01 := lengthOfLongestSubstring(test.s)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, sum01)
		}
	}
}
