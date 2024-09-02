package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/minimum-size-subarray-sum/description/

func minSubArrayLen0902(target int, nums []int) int {
	ans, sum, left := len(nums)+1, 0, 0
	for right, num := range nums {
		sum += num
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}

	return ans
}

func TestMinSubArrayLen(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{2, 3, 1, 2, 4, 3}, target: 7, expect: 2},
		{array: []int{1, 1, 1, 1, 1, 1, 1, 1}, target: 11, expect: 0},
	}

	for _, test := range twoSumTest {
		get := minSubArrayLen0902(test.target, test.array)
		if get != test.expect {
			t.Fatalf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
