package _0_basic

import (
	"fmt"
	"math/bits"
	"testing"
)

// https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/description/

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	if n|k > n {
		return -1
	}
	tmp := n ^ k

	var res int
	for tmp > 0 {
		if tmp&1 == 1 {
			res++
		}
		tmp >>= 1
	}

	return res
}

func minChanges01(n int, k int) int {
	if n|k > n {
		return -1
	}

	return bits.OnesCount(uint(n ^ k))
}

func TestMinChanges(t *testing.T) {
	fmt.Println(minChanges01(13, 4))
}
