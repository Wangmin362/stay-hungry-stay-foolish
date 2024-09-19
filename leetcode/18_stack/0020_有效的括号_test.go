package _1_array

import (
	"container/list"
	"testing"
)

// 题目：https://leetcode.cn/problems/valid-parentheses/description/

// 解题思路：做括号入栈，右括号出栈
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

func isValid01(s string) bool {
	mapping := map[uint8]uint8{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := list.New()
	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack.PushBack(uint8(c))
		} else {
			if stack.Len() <= 0 {
				return false
			}
			cs := stack.Remove(stack.Back()).(uint8)
			if pair, ok := mapping[uint8(c)]; !ok || pair != cs {
				return false
			}
		}
	}
	if stack.Len() > 0 {
		return false
	}

	return true
}

func isValid03(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	if len(s) == 0 {
		return true
	}

	c := map[byte]byte{'}': '{', ']': '[', ')': '('}
	var arr []byte
	for _, char := range s {
		mapping, ok := c[byte(char)]
		if !ok {
			arr = append(arr, byte(char))
		} else {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] != mapping {
				return false
			}
			arr = arr[:len(arr)-1]
		}
	}

	if len(arr) > 0 {
		return false
	}

	return true
}

func TestReverseWords(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect bool
	}{
		{s: "[", expect: false},
		{s: "]", expect: false},
		{s: "[]", expect: true},
		{s: "[)", expect: false},
		{s: "[][]{}", expect: true},
		{s: "[][]{}", expect: true},
		{s: "[{}{}[][{}]]", expect: true},
	}

	for _, test := range teatdata {
		get := isValid03(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
