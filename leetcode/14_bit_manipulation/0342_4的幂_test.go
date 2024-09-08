package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/power-of-four/description/

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}

	tmp := 4
	for i := 1; i < n; i++ {
		if tmp == n {
			return true
		} else if tmp > n {
			return false
		}
		tmp *= 4
	}
	return false
}

func isPowerOfFour01(n int) bool {
	return n&(n-1) == 0 && (n%3 == 1)
}

func TestIsPowerOfFour(t *testing.T) {
	fmt.Println(isPowerOfFour01(64))
}
