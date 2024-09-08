package _0_basic

import (
	"fmt"
	"math/bits"
	"testing"
)

// https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
func TestHammingDistance(t *testing.T) {
	fmt.Println(hammingDistance(4, 5))
}
