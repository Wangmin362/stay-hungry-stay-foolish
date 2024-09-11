package _0_basic

import (
	"strings"
	"testing"
)

// https://leetcode.cn/problems/reverse-words-in-a-string-iii/description/?envType=problem-list-v2&envId=two-pointers&difficulty=EASY

func reverseWords(s string) string {
	reverse := func(str string) string {
		s := []byte(str)
		left, right := 0, len(s)-1
		for left < right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
		return string(s)
	}
	sp := strings.Split(s, " ")
	for idx, ss := range sp {
		sp[idx] = reverse(ss)
	}

	return strings.Join(sp, " ")
}

func TestReverseWords(t *testing.T) {
	var testData = []struct {
		s    string
		want string
	}{
		{s: "Let's take LeetCode contest", want: "s'teL ekat edoCteeL tsetnoc"},
		{s: "Mr Ding", want: "rM gniD"},
	}

	for _, tt := range testData {
		get := reverseWords(tt.s)
		if get != tt.want {
			t.Fatalf("s:%v, want:%v, get:%v", tt.s, tt.want, get)
		}
	}
}
