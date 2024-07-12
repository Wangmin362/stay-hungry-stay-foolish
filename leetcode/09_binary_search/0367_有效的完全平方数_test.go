package _9_binary_search

import (
	"testing"
)

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left := 0
	right := num
	for left <= right {
		mid := left + (right-left)>>1
		mul := mid * mid // 平方数
		if mul == num {
			return true
		} else if mul < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func isPerfectSquare02(num int) bool {
	if num <= 1 {
		return true
	}
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)>>1
		pow := mid * mid
		if pow == num {
			return true
		} else if pow > num {
			right = mid - 1
		} else if pow < num {
			left = mid + 1
		}
	}
	return false
}

func TestIsPerfectSquare(t *testing.T) {
	var twoSumTest = []struct {
		target int
		expect bool
	}{
		{target: 0, expect: true},
		{target: 1, expect: true},
		{target: 2, expect: false},
		{target: 3, expect: false},
		{target: 4, expect: true},
		{target: 8, expect: false},
		{target: 9, expect: true},
		{target: 10, expect: false},
		{target: 16, expect: true},
		{target: 25, expect: true},
		{target: 49, expect: true},
	}

	for _, test := range twoSumTest {
		get := isPerfectSquare02(test.target)
		if get != test.expect {
			t.Fatalf("target:%v, expect:%v, get:%v", test.target, test.expect, get)
		}
	}
}
