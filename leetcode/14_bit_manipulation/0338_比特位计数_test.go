package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/counting-bits/description/

func countBits(n int) []int {
	cnt := func(n int) int {
		var res int
		for n > 0 {
			if n&1 == 1 {
				res++
			}
			n >>= 1
		}
		return res
	}

	res := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		res = append(res, cnt(i))
	}

	return res
}

func TestCountBits(t *testing.T) {
	fmt.Println(countBits(2))
}
