package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/valid-palindrome/description/?envType=study-plan-v2&envId=top-interview-150

// 使用碰撞指针
func isPalindrome(s string) bool {
	isAlpha := func(c byte) bool {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			return true
		}
		return false
	}
	isNumber := func(c byte) bool {
		if c >= '0' && c <= '9' {
			return true
		}
		return false
	}
	isAlphaNumber := func(c byte) bool {
		return isAlpha(c) || isNumber(c)
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphaNumber(s[left]) {
			left++
		}
		for left < right && !isAlphaNumber(s[right]) {
			right--
		}
		if left >= right {
			break
		}
		if s[left] != s[right] {
			if (isAlpha(s[left]) && isNumber(s[right])) || (isAlpha(s[right]) && isNumber(s[left])) {
				return false
			}
			if s[left]-s[right] != 'a'-'A' && s[right]-s[left] != 'a'-'A' {
				return false
			}
		}

		left++
		right--
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	var testdata = []struct {
		s    string
		want bool
	}{
		//{s: "A man, a plan, a canal: Panama", want: true},
		{s: "0P", want: false},
	}

	for _, tt := range testdata {
		get := isPalindrome(tt.s)
		if get != tt.want {
			t.Fatalf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
