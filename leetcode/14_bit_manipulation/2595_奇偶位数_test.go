package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/number-of-even-and-odd-bits/description/

func evenOddBit(n int) []int {
	even, odd, idx := 0, 0, 0
	for n > 0 {
		if n&1 == 1 {
			if idx&1 == 0 { // 偶数坐标
				even++
			} else { // 奇数坐标
				odd++
			}
		}
		n >>= 1
		idx++
	}
	return []int{even, odd}
}

func TestEvenOddBit(t *testing.T) {
	fmt.Println(evenOddBit(17))
}
