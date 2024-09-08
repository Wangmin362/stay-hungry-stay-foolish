package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/is-unique-lcci/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func isUnique(astr string) bool {
	// 每个bit代表一个英文字母
	cache := 0
	for _, c := range astr {
		t := 1 << (c - 'a')
		if cache&t != 0 {
			return false
		}
		cache |= t
	}
	return true
}

func TestIsUnique(t *testing.T) {
	fmt.Println(isUnique("leetcode"))
	fmt.Println(isUnique("abc"))
}
