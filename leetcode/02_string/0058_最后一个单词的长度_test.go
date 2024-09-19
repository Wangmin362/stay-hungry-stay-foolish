package _1_array

import (
	"strings"
)

// https://leetcode.cn/problems/length-of-last-word/description/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLastWord01(s string) int {
	s = strings.Trim(s, " ")
	sp := strings.Split(s, " ")

	return len(sp[len(sp)-1])
}

func lengthOfLastWord02(s string) int {
	idx := len(s) - 1
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}
	if idx < 0 {
		return 0
	}
	first := idx

	for idx >= 0 && s[idx] != ' ' {
		idx--
	}
	return first - idx
}
