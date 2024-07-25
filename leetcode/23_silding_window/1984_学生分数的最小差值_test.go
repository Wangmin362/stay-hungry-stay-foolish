package _1_array

import (
	"slices"
	"testing"
)

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/

func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)
	if len(nums) <= 1 {
		return 0
	}
	if len(nums) <= k {
		return nums[len(nums)-1] - nums[0]
	}
	left, right := 0, k-1
	minDiff := nums[len(nums)-1]
	for right < len(nums) {
		diff := nums[right] - nums[left]
		minDiff = min(minDiff, diff)
		left++
		right++
	}
	return minDiff
}

func TestMinimumDifference(t *testing.T) {
	testdata := []struct {
		nums   []int
		k      int
		expect int
	}{
		{nums: []int{9, 4, 1, 7}, k: 2, expect: 2},
	}

	for _, test := range testdata {
		get := minimumDifference(test.nums, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
