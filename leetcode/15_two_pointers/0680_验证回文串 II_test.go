package _0_basic

import "testing"

// https://leetcode.cn/problems/valid-palindrome-ii/?envType=problem-list-v2&envId=two-pointers&difficulty=EASY

// 从两边向中间走，允许一次遇到不等的字符，第二次遇到不等就认为无效
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	hasDelete := false
	valid := true
	for left < right {
		if s[left] != s[right] {
			if !hasDelete {
				hasDelete = true
				right++ // 右边不移动，相当于删除左边
			} else {
				valid = false
				break
			}
		}
		left++
		right--
	}
	if valid {
		return true
	}

	hasDelete = false
	left, right = 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if !hasDelete {
				hasDelete = true
				left-- // 左边不移动，相当于删除右边
			} else {
				return false
			}
		}
		left++
		right--
	}

	return true
}

func TestName(t *testing.T) {
	var testData = []struct {
		s    string
		want bool
	}{
		//{s: "aba", want: true},
		//{s: "abca", want: true},
		{s: "abc", want: false},
		//{s: "deeee", want: true},
		//{s: "cbbcc", want: true},
	}
	for _, tt := range testData {
		get := validPalindrome(tt.s)
		for get != tt.want {
			t.Fatalf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
