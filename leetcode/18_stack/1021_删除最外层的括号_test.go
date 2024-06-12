package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-outermost-parentheses/

func removeOuterParentheses(s string) string {
	var stack rune
	var resStack []rune
	var res string
	left := 0
	for _, char := range s {
		if stack == 0 {
			stack = char
		} else {
			if len(resStack) > 0 && resStack[len(resStack)-1] == ')' && char == ')' && left == 0 {
				res += string(resStack)
				resStack = []rune{}
				stack = 0
				left = 0
			} else if len(resStack) == 0 && char == ')' {
				stack = 0
				left = 0
			} else {
				// (()())(())(()(()))
				if char == '(' {
					left++
				} else {
					left--
				}
				resStack = append(resStack, char)
			}
		}
	}
	return res
}

func TestRemoveOuterParentheses(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect string
	}{
		{s: "()()()()(())", expect: "()"},
		{s: "(()())(())(()(()))", expect: "()()()()(())"},
		{s: "(()())(())", expect: "()()()"},
		{s: "()()", expect: ""},
		{s: "", expect: ""},
	}

	for _, test := range teatdata {
		get := removeOuterParentheses(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
