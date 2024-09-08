package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/number-of-1-bits/description/

func hammingWeight(n int) int {
	var res int
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}
	return res
}
func TestHammingWeight(t *testing.T) {
	fmt.Println(hammingWeight(17))
}
