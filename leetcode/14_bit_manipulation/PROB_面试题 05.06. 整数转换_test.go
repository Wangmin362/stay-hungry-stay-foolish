package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/convert-integer-lcci/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=EASY

func convertInteger(A int, B int) int {
	xor := uint32(A ^ B)
	var res int
	for xor != 0 {
		if xor&0x1 == 0x1 {
			res++
		}
		xor >>= 1
	}
	return res
}

func TestConvertInteger(t *testing.T) {
	fmt.Println(convertInteger(826966453, -729934991))

}
