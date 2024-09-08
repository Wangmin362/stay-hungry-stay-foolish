package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/reverse-bits/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func reverseBits(num uint32) uint32 {
	var res uint32
	idx := 31
	for num > 0 {
		curr := num & 0x1
		res += curr * (1 << idx)
		idx--
		num >>= 1
	}
	return res
}

func TestReverseBits(t *testing.T) {
	fmt.Println(reverseBits(43261596))
}
