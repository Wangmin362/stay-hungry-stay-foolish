package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/binary-number-with-alternating-bits/description/

func hasAlternatingBits(n int) bool {
	prev := -1
	for n > 0 {
		curr := n & 1
		if prev == -1 {
			prev = curr
		} else if prev == curr {
			return false
		} else {
			prev = curr
		}
		n >>= 1
	}
	return true
}

func TestHasAlternatingBits(t *testing.T) {
	fmt.Println(hasAlternatingBits(5))
}
