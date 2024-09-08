package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sum-of-two-integers/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=MEDIUM

func getSum(a int, b int) int {
	for b != 0 {
		c := uint32(a&b) << 1
		a ^= b
		b = int(c)
	}
	return a
}

func TestGetSum(t *testing.T) {
	fmt.Println(getSum(2, 3))
}
