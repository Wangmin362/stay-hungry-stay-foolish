package _9_binary_search

import (
	"testing"
)

// 解题思路: 计算某个数的算数平方根，其实就是 x * x = target，只需要找到这个x即可，找到第一次x * x <= target的x就是我们要找的那个值。

// target = 9  x=4 x^2=16 > 9
// x=2 x^2=4 < 9

func mySqrt(x int) int {
	target := x

	left := 0
	right := x

	res := -1
	for left <= right {
		mul := x * x
		if mul == target {
			return x
		} else if mul > target {
			right = x - 1
			x >>= 1 // 继续除2
		} else { // 说明 mul < target
			// 如果第一次发现小于target
			res = x // 记录下最后一次小于target的数字
			left = x + 1
			x++ // 向右边移动一次
		}
	}

	return res
}

func mySqrt01(x int) int {
	if x < 2 {
		return x
	}

	left := 0
	right := x
	ans := 0
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func TestMySqrt(t *testing.T) {
	var twoSumTest = []struct {
		target int
		expect int
	}{
		{target: 0, expect: 0},
		{target: 1, expect: 1},
		{target: 2, expect: 1},
		{target: 3, expect: 1},
		{target: 4, expect: 2},
		{target: 8, expect: 2},
		{target: 9, expect: 3},
	}

	for _, test := range twoSumTest {
		get := mySqrt01(test.target)
		if get != test.expect {
			t.Fatalf("target:%v, expect:%v, get:%v", test.target, test.expect, get)
		}
	}
}
