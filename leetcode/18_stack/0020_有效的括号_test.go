package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/valid-parentheses/description/

// 接替思路：做括号入栈，右括号出栈
func isValid(s string) bool {
	var stack []rune
	for _, sChar := range s {
		if sChar == '{' || sChar == '[' || sChar == '(' {
			stack = append(stack, sChar)
		} else {
			if len(stack) <= 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if sChar == ']' && last != '[' {
				return false
			} else if sChar == '}' && last != '{' {
				return false
			} else if sChar == ')' && last != '(' {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}
func TestReverseWords(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect bool
	}{
		{s: "[]", expect: true},
		{s: "[)", expect: false},
		{s: "[][]{}", expect: true},
		{s: "[][]{}", expect: true},
		{s: "[{}{}[][{}]]", expect: true},
	}

	for _, test := range teatdata {
		get := isValid(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
