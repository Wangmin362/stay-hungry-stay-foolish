package _1_array

import (
	"testing"
)

// 只能使用快慢指着， 对撞指针会移动元素顺序

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func TestRemoveDuplicates(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{3, 3, 3}, expect: []int{3}},
		{array: []int{3, 3, 3, 4}, expect: []int{3, 4}},
		{array: []int{3, 3, 3, 4, 5, 6}, expect: []int{3, 4, 5, 6}},
		{array: []int{3}, expect: []int{3}},
		{array: []int{3, 3, 3, 3, 3}, expect: []int{3}},
		{array: []int{3, 3}, expect: []int{3}},
		{array: []int{2, 2, 3}, expect: []int{2, 3}},
		{array: []int{2, 2, 2, 2, 2, 2, 3}, expect: []int{2, 3}},
		{array: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}, expect: []int{2, 3}},
		{array: []int{0, 1, 2, 2, 3, 4}, expect: []int{0, 1, 2, 3, 4}},
	}

	for _, test := range twoSumTest {
		get := removeDuplicates(test.array)
		if get != len(test.expect) {
			t.Errorf("expect:%v, get:%v", len(test.expect), get)
		}

		for i := 0; i < get; i++ {
			if test.array[i] != test.expect[i] {
				t.Errorf("expect:%v, get:%v", test.expect, test.array)
			}
		}
	}
}
