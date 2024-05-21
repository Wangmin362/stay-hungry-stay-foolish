package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/longest-valid-parentheses/description/

func longestValidParentheses(s string) int {
	var stack []rune
	var res string
	for _, sChar := range s {
		if len(stack) == 0 {
			stack = append(stack, sChar)
		} else {
			last := stack[len(stack)-1]
			if sChar == ')' && last == '(' {
				if len(stack) > 2 {

				}
				res += "()"
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, sChar)
			}
		}
	}

	return len(res)
}

func TestLongestValidParentheses(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect int
	}{
		{s: "()", expect: 2},
		{s: "((", expect: 0},
		{s: "(()", expect: 2},
		{s: "(())", expect: 4},
		{s: "()(()", expect: 2},
		{s: "()(()()", expect: 4},
	}

	for _, test := range teatdata {
		get := longestValidParentheses(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
