package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/xor-operation-in-an-array/description/

func xorOperation(n int, start int) int {
	var res int
	for i := 0; i < n; i++ {
		num := start + 2*i
		if i == 0 {
			res = num
		} else {
			res ^= num
		}
	}
	return res
}

func TestXorOperation(t *testing.T) {
	fmt.Println(xorOperation(5, 0))
}
