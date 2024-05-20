package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/repeated-substring-pattern/description/

// 使用切片
func repeatedSubstringPattern01(s string) bool {
	mid := len(s) >> 1
	for sp := 1; sp <= mid; sp++ {
		idx := 0
		isValid := true
		for idx < len(s) {
			s1Idx := idx
			s2Idx := s1Idx + sp
			s3Idx := s2Idx + sp

			if s3Idx > len(s) {
				isValid = false
				break
			}

			s1 := s[s1Idx:s2Idx]
			s2 := s[s2Idx:s3Idx]
			if s1 != s2 {
				isValid = false
				break
			}
			if s3Idx == len(s) {
				break
			}

			idx = s2Idx
		}
		if isValid {
			return true
		}
	}

	return false
}

func TestRepeatedSubstringPattern(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect bool
	}{
		{s: "abab", expect: true},
		{s: "aba", expect: false},
		{s: "aaa", expect: true},
		{s: "abcabc", expect: true},
	}

	for _, test := range teatdata {
		get := repeatedSubstringPattern01(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
