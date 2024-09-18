package _1_array

import (
	"reflect"
	"testing"
)

// 题目要求：不能使用额外的空间，并且使用O(n)的时间复杂度
// 解题思路：使用两个指针的方式来解题。有两种方式，一种是快慢指针，一种是对撞指针
// 快慢指针：不会改变元素的相对顺序
// 对撞指针：会改变元素的相对顺序

// 快慢指针
func removeElement02(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 对撞指针
func removeElement03(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left >= right {
			return left
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	return left
}

func TestRemoveElement(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{3, 2, 2, 3}, target: 3, expect: 2},
		{array: []int{0, 1, 2, 2, 3, 0, 4, 2}, target: 2, expect: 5},
		{array: []int{2}, target: 3, expect: 1},
	}

	for _, test := range twoSumTest {
		total := removeElement02(test.array, test.target)
		if !reflect.DeepEqual(total, test.expect) {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, total)
		}
	}
}
