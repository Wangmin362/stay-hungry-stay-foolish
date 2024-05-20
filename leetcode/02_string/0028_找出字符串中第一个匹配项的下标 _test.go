package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

func strStr(haystack string, needle string) int {

	for idx := 0; idx < len(haystack)-len(needle)+1; idx++ {
		if haystack[idx:idx+len(needle)] == needle {
			return idx
		}
	}

	return -1
}

func TestStrStr(t *testing.T) {
	var teatdata = []struct {
		haystack string
		needle   string
		expect   int
	}{
		{haystack: "abcedfeg", needle: "ced", expect: 2},
	}

	for _, test := range teatdata {
		get := strStr(test.haystack, test.needle)
		if get != test.expect {
			t.Errorf("haystack: %v, needle:%v, expect:%v, get:%v", test.haystack, test.needle, test.expect, get)
		}
	}
}
