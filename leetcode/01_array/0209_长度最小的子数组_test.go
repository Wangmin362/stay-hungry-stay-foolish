package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/minimum-size-subarray-sum/description/
// 接替思路：使用快慢指针，计算快慢指针之间所有数字总和。若数字总和小于target，就移动快指针。若数字总和大于target，就移动慢指针。同时，
// 动态更新满足的长度

func minSubArrayLen(target int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	left := 0
	right := 0
	sum := nums[0]
	minLen := len(nums) + 1
	for right < len(nums) && left <= right {
		if sum >= target {
			currLen := right - left + 1
			if minLen > currLen {
				minLen = currLen
			}

			sum -= nums[left]
			left++ // 移动左指针，缩小范围，减小总和
		} else {
			right++ // 移动右指针，扩大范围，增加总和
			if right < len(nums) {
				sum += nums[right]
			}
		}
	}

	if minLen == len(nums)+1 { // 说明不存在
		return 0
	}

	return minLen
}

func TestMinSubArrayLen(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{2, 3, 1, 2, 4, 3}, target: 7, expect: 2},
	}

	for _, test := range twoSumTest {
		get := minSubArrayLen(test.target, test.array)
		if get != test.expect {
			t.Fatalf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
