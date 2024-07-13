package _1_array

import (
	"reflect"
	"testing"
)

// 题目要求：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 解题思路：使用快慢指针，快指针用于遍历元素，满指针用于指向非零的位置

func moveZeroes(nums []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

func moveZeroes02(nums []int) {
	if len(nums) <= 1 {
		return
	}
	slow, fast := 0, 0
	for slow < len(nums) && fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

func TestMoveZeroes(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{0, 1, 0, 3, 12}, expect: []int{1, 3, 12, 0, 0}},
		{array: []int{0, 1, 0, 0, 0}, expect: []int{1, 0, 0, 0, 0}},
		{array: []int{0, 1}, expect: []int{1, 0}},
		{array: []int{1, 0}, expect: []int{1, 0}},
		{array: []int{1}, expect: []int{1}},
		{array: []int{0}, expect: []int{0}},
	}

	for _, test := range twoSumTest {
		moveZeroes02(test.array)
		if !reflect.DeepEqual(test.array, test.expect) {
			t.Errorf("expect:%v, get:%v", test.expect, test.array)
		}
	}
}
