package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/reverse-vowels-of-a-string/description/?envType=problem-list-v2&envId=two-pointers&difficulty=EASY

func reverseVowels(s string) string {
	isVowel := func(c byte) bool {
		if c == 'a' || c == 'A' ||
			c == 'e' || c == 'E' ||
			c == 'i' || c == 'I' ||
			c == 'o' || c == 'O' ||
			c == 'u' || c == 'U' {
			return true
		}
		return false
	}

	left, right := 0, len(s)-1
	sp := []byte(s)
	for left < right {
		for left < right && !isVowel(sp[left]) {
			left++
		}
		for left < right && !isVowel(sp[right]) {
			right--
		}
		if left >= right {
			break
		}
		sp[left], sp[right] = sp[right], sp[left]
		left++
		right--
	}
	return string(sp)
}

func TestReverseVowels(t *testing.T) {
	var testData = []struct {
		s      string
		expect string
	}{
		{s: "hello", expect: "holle"},
		{s: "leetcode", expect: "leotcede"},
	}

	for _, test := range testData {
		get := reverseVowels(test.s)
		if get != test.expect {
			t.Errorf("s:%v, expect:%v, get:%v", test.s, test.expect, get)
		}
	}

}
