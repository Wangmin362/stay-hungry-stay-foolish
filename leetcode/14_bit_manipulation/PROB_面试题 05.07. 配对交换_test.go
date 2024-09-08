package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/exchange-lcci/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func exchangeBits(num int) int {
	odd := uint32(num) & 0x5555_5555
	even := uint32(num) & 0xaaaa_aaaa
	return int(odd<<1 | even>>1)
}
func TestExchangeBits(t *testing.T) {
	fmt.Println(exchangeBits(3))
}
