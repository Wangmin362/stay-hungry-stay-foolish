package _1_array

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/numbers-with-same-consecutive-differences/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func numsSameConsecDiff(n int, k int) []int {
	cnt := int(math.Pow(float64(10), float64(n)))
	begin := int(math.Pow(float64(10), float64(n-1)))

	isValidNum := func(num int) bool {
		prev := num % 10
		x := num
		x /= 10
		for x > 0 {
			if x%10-k != prev && x%10+k != prev {
				return false
			}
			prev = x % 10
			x /= 10
		}

		return true
	}

	var res []int
	for i := begin; i < cnt; i++ {
		if isValidNum(i) {
			res = append(res, i)
		}
	}

	return res
}
func TestNumsSameConsecDiff(t *testing.T) {
	fmt.Println(numsSameConsecDiff(3, 7))
}
