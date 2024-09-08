package _0_basic

import (
	"fmt"
	"math/bits"
	"testing"
)

// https://leetcode.cn/problems/minimum-bit-flips-to-convert-number/description/

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
func TestMinBitFlips(t *testing.T) {
	fmt.Println(minBitFlips(5, 0))
}
