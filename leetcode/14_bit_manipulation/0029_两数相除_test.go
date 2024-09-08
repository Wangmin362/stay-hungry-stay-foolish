package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/divide-two-integers/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=MEDIUM

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}

	var res int
	adder := divisor
	if divisor < 0 {
		adder = -adder
	}
	summer := dividend
	if dividend < 0 {
		summer = -summer
	}
	for adder <= summer {
		res++
		if divisor > 0 {
			adder += divisor
		} else {
			adder -= divisor
		}
	}

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		return -res
	}

	return res
}

func TestDivide(t *testing.T) {
	fmt.Println(divide(2147483647, 2))
}
