package _1_array

import (
	"reflect"
	"testing"
)

// 题目要求：不能使用额外的空间，并且使用O(n)的时间复杂度
// 接替思路：使用两个指针的方式来解题。有两种方式，一种是快慢指针，一种是对撞指针。根据题目的意思元素的相对顺序是不能改变的，因此只能使用
// 快慢指针，对撞指针会改变元素的相对顺序

func removeElement01(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val { // 当前位置不需要移除
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func TestRemoveElement(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{3, 2, 2, 3}, target: 3, expect: 2},
		{array: []int{0, 1, 2, 2, 3, 0, 4, 2}, target: 2, expect: 5},
	}

	for _, test := range twoSumTest {
		total := removeElement01(test.array, test.target)
		if !reflect.DeepEqual(total, test.expect) {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, total)
		}
	}
}
