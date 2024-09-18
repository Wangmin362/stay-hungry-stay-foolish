package _1_array

import (
	"testing"
)

func canJump(nums []int) bool {
	iMax := 0 // 初始化能跳到最远的位置就是0
	idx, jump := 0, 0
	// 遍历数组，如果当前最远的在当前位置的后面，说明当前位置时可以跳到的，因此
	// 只要当前位置加上可以跳到的最远位置大于记录的最远位置，就更新最远位置，最后比较
	// 最远位置和最后一个位置
	for idx, jump = range nums {
		if iMax >= idx && idx+jump > iMax {
			iMax = idx + jump
		}
	}

	return iMax >= idx
}

func canJump02(nums []int) bool {
	mx := 0 // 记录可以跳到的最远的位置
	for i, jump := range nums {
		if i > mx { // 说明当前位置根本跳不到
			return false
		}
		mx = max(mx, i+jump)
	}
	return true
}

func TestCanJump(t *testing.T) {
	twoSumTest := []struct {
		array  []int
		expect bool
	}{
		{array: []int{2, 3, 1, 1, 4}, expect: true},
		{array: []int{3, 2, 1, 0, 4}, expect: false},
	}

	for _, test := range twoSumTest {
		get := canJump02(test.array)
		if get != test.expect {
			t.Errorf("target:%v, expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
