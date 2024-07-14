package _1_array

import (
	"container/list"
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/

// 接替思路：使用栈来解决这个问题，遇到相同的
func removeDuplicates(s string) string {
	var res []rune
	for _, ch := range s {
		if len(res) == 0 {
			res = append(res, ch)
		} else {
			if ch == res[len(res)-1] {
				res = res[:len(res)-1]
			} else {
				res = append(res, ch)
			}
		}
	}
	return string(res)
}

func removeDuplicates01(s string) string {
	stack := list.New()
	for _, c := range s {
		if stack.Len() <= 0 {
			stack.PushBack(c)
		} else {
			sc := stack.Back().Value.(rune)
			if sc == c {
				stack.Remove(stack.Back())
			} else {
				stack.PushBack(c)
			}
		}
	}
	var str []rune
	for stack.Len() > 0 {
		str = append(str, stack.Remove(stack.Front()).(rune))
	}

	return string(str)
}

func TestRemoveDuplicates(t *testing.T) {
	var teatdata = []struct {
		s      string
		expect string
	}{
		{s: "abbaca", expect: "ca"},
		{s: "abbbaca", expect: "abaca"},
		{s: "bbbbaca", expect: "aca"},
		{s: "babbbaca", expect: "babaca"},
	}

	for _, test := range teatdata {
		get := removeDuplicates01(test.s)
		if get != test.expect {
			t.Errorf("s: %v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}
}
