package _1_array

import (
	"testing"
)

func canJumpII(nums []int) int {
	iMax := 0 // 初始化能跳到最远的位置就是0
	idx, jump := 0, 0
	// 遍历数组，如果当前最远的在当前位置的后面，说明当前位置时可以跳到的，因此
	// 只要当前位置加上可以跳到的最远位置大于记录的最远位置，就更新最远位置，最后比较
	// 最远位置和最后一个位置
	for idx, jump = range nums {
		if iMax >= idx && idx+jump > iMax {
			iMax = idx + jump
		}
		if iMax >= len(nums)-1 { // 只要能跳的最大位置超过最后的位置，那么就可以达到，此时直接返回
			return idx + 1
		}
	}

	return idx + 1
}

func TestCanJumpII(t *testing.T) {
	twoSumTest := []struct {
		array  []int
		expect int
	}{
		{array: []int{2, 3, 1, 1, 4}, expect: 2},
		{array: []int{3, 2, 1, 0, 4}, expect: 0},
	}

	for _, test := range twoSumTest {
		get := canJumpII(test.array)
		if get != test.expect {
			t.Errorf("target:%v, expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
