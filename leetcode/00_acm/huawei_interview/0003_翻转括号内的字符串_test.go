package huawei_interview

import (
	"strings"
	"testing"
)

// 给出一个字符串（仅含有小写英文字母和括号）。
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
// 输入：(abcd)   输出：dcba
// 输入：(u(love)i) 输出：iloveu
// 输入：(ed(et(oc))le) 输出：leetcode

// 解题思路：使用栈
func reverseStr(str string) string {
	reverse := func(str string) string {
		bytes := []byte(str)
		left, right := 0, len(bytes)-1
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
		return string(bytes)
	}
	stack := make([]string, 0, len(str))
	left, right := 0, 0
	for right = 0; right < len(str); right++ {
		if str[right] == '(' {
			s := str[left:right]
			if len(s) != 0 {
				stack = append(stack, s) //入栈字符串
			}

			stack = append(stack, str[right:right+1]) // 入栈左括号
			left = right + 1
		} else if str[right] == ')' { // 找到所有的左括号
			s := str[left:right]
			for stack[len(stack)-1] != "(" {
				s = stack[len(stack)-1] + s
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // 去除左括号

			s = reverse(s)
			stack = append(stack, s)
			left = right + 1
		}
	}
	if left < len(str) {
		s := str[left:right]
		stack = append(stack, s)
	}

	return strings.Join(stack, "")
}

func TestReverseStr(t *testing.T) {
	var testdata = []struct {
		str  string
		want string
	}{
		{str: "(abcd)", want: "dcba"},
		{str: "(u(love)i)", want: "iloveu"},
		{str: "(ed(et(oc))el)", want: "leetcode"},
		{str: "a(bcdefghijkl(mno)p)q", want: "apmnolkjihgfedcbq"},
	}
	for _, tt := range testdata {
		get := reverseStr(tt.str)
		if get != tt.want {
			t.Fatalf("str:%v, want:%v, get:%v", tt.str, tt.want, get)
		}
	}
}
