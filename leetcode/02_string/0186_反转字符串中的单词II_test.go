package _1_array

import (
	"fmt"
	"testing"
)

func reverseWords(s []byte) {
	reverse := func(start, end int) {
		for start < end {
			s[start], s[end] = s[end], s[start]
			start++
			end--
		}
	}

	reverse(0, len(s)-1)
	start, end := 0, 0
	for end < len(s) {
		for start < len(s) && s[start] == ' ' { // 找单词的第一个空格
			start++
		}

		for end < len(s) && s[end] != ' ' { // 找单词的最后一个空格
			end++
		}
		if start < len(s) {
			reverse(start, end-1)
			start = end
			end += 1
		}
	}

	if start < len(s) {
		reverse(start, end-1)
	}
}

func TestReverseWordsII(t *testing.T) {
	word := []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}
	fmt.Println(string(word))
	reverseWords(word)
	fmt.Println(string(word))
}
