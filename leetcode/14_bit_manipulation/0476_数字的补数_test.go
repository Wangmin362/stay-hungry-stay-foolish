package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/number-complement/description/

func findComplement(num int) int {
	x := uint32(num)
	bLen := 0
	for num > 0 {
		bLen++
		num >>= 1
	}
	x = ^x
	x &= 1<<bLen - 1

	return int(x)
}

func TestFindComplement(t *testing.T) {
	fmt.Println(findComplement(5))
}
