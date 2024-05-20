package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/reverse-string-ii/description/

func reverseStrii(s string, k int) string {
	byteS := []byte(s)

	reverserFun := func(startIdx, endIdx int) {
		for startIdx < endIdx {
			byteS[startIdx], byteS[endIdx] = byteS[endIdx], byteS[startIdx]
			startIdx++
			endIdx--
		}
	}

	startIdx := 0
	for startIdx < len(byteS) {
		reverseIdx := startIdx + k  // [0, k)
		doubleIdx := startIdx + 2*k // [0,2*k)

		if doubleIdx <= len(byteS) { // 反转前k个字符
			reverserFun(startIdx, reverseIdx-1)
			startIdx = doubleIdx
		} else { // 剩余没有2k个字符
			cnt := len(byteS) - startIdx
			if cnt < k { // 反转剩余全部字符
				reverserFun(startIdx, len(byteS)-1)
			} else { // 反转k个字符
				reverserFun(startIdx, reverseIdx-1)
			}
			break
		}
	}

	return string(byteS)
}

func TestReverseStrii(t *testing.T) {
	var teatdata = []struct {
		s      string
		k      int
		expect string
	}{
		{s: "abcdefg", k: 2, expect: "bacdfeg"},
		{s: "abcd", k: 2, expect: "bacd"},
	}

	for _, test := range teatdata {
		get := reverseStrii(test.s, test.k)
		if get != test.expect {
			t.Errorf("expect:%v, get:%v", test.expect, get)
		}
	}
}
