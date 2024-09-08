package _0_basic

import "testing"

// https://leetcode.cn/problems/palindrome-permutation-lcci/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func canPermutePalindrome(s string) bool {
	c := make(map[rune]int)
	for _, cc := range s {
		c[cc]++
	}
	if len(s)%2 == 0 {
		for _, cnt := range c {
			if cnt != 2 {
				return false
			}
		}
		return true
	} else {
		one := 0
		for _, cnt := range c {
			if cnt == 1 {
				one++
			}
		}
		if one == 1 {
			return true
		}
		return false
	}
}

func TestCanPermutePalindrome(t *testing.T) {
	canPermutePalindrome("abdg")
}
