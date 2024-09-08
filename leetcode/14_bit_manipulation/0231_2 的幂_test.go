package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/power-of-two/description/

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func TestIsPowerOfTwo(t *testing.T) {
	fmt.Println(isPowerOfTwo(16))
}
