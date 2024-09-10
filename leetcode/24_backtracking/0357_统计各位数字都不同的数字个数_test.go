package _1_array

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/count-numbers-with-unique-digits/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func countNumbersWithUniqueDigits(n int) int {
	isValidNum := func(num int) bool {
		if num < 10 {
			return true
		}

		ge := num % 10
		xnum := num
		xnum /= 10
		for xnum > 0 {
			mod := xnum % 10
			if mod != ge {
				return true
			}
			xnum /= 10
		}

		return false
	}

	n = int(math.Pow(float64(10), float64(n)))

	var res int
	for i := 0; i < n; i++ {
		if isValidNum(i) {
			res++
		}
	}
	return res
}

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	fmt.Println(countNumbersWithUniqueDigits(3))
}
