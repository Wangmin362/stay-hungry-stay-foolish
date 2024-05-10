package _9_binary_search

import (
	"testing"
)

// 题目：给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。

// 分析，c = a^2 + b^2， 一定是一个小数字的平方加上一个大数字的平方 left=0, right=c, 左边界0也是符合条件的，此时如果符合的话，说明c整好是
// 一个完全平方数。 left^2 + right^2 = c

// 0  => 0^2 + 0^2
// 1  => 0^2 + 1^2
// 2  => 1^2 + 1^2
// 4  => 0^2 + 2^2
// 5  => 1^2 + 2^2
// 8  => 2^2 + 2^2
// 9  => 0^2 + 3^2
// 10  => 1^2 + 3^2

func judgeSquareSum(c int) bool {
	left := 0
	right := c
	for left <= right {
		mid := left + (right-left)>>1

		mul := mid * mid
		if mul == c { // 正好是一个完全平方数
			return true
		} else if mul > c { // 中位数的平方大于target，那么搜索域一定在中位数左边
			right = mid - 1
		} else {
			sumMul := left*left + mul
			if sumMul == c {
				return true
			} else if sumMul > c {
				right = right - 1 // 这里有边界应该缩小一个
			} else {
				left = left + 1
			}
		}

	}

	return false
}

func TestJudgeSquareSum(t *testing.T) {
	var twoSumTest = []struct {
		target int
		expect bool
	}{
		//{target: 0, expect: true},
		//{target: 1, expect: true},
		//{target: 2, expect: true},
		//{target: 3, expect: false},
		//{target: 4, expect: true},
		{target: 5, expect: true},
		{target: 8, expect: true},
		{target: 9, expect: true},
		{target: 10, expect: true},
		{target: 16, expect: true},
		{target: 25, expect: true},
		{target: 49, expect: true},
	}

	for _, test := range twoSumTest {
		get := judgeSquareSum(test.target)
		if get != test.expect {
			t.Fatalf("target:%v, expect:%v, get:%v", test.target, test.expect, get)
		}
	}
}
