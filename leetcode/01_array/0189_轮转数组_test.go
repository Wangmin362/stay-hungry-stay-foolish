package _1_array

import (
	"reflect"
	"testing"
)

func rotate01(nums []int, k int) {
	if k <= 0 {
		return
	}
	k = k % len(nums)
	n1, n2 := make([]int, len(nums)-k), make([]int, k)
	for idx := range nums {
		if idx < len(nums)-k {
			n1[idx] = nums[idx]
		} else {
			n2[idx-(len(nums)-k)] = nums[idx]
		}
	}
	for idx := range nums {
		if idx < k {
			nums[idx] = n2[idx]
		} else {
			nums[idx] = n1[idx-k]
		}
	}
}

// 先反转，再反转
func rotate02(nums []int, k int) {
	if k <= 0 {
		return
	}

	k = k % len(nums)

	reverse := func(nums []int, begin, end int) {
		for begin < end {
			nums[begin], nums[end] = nums[end], nums[begin]
			begin++
			end--
		}
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func TestRotate(t *testing.T) {
	twoSumTest := []struct {
		array  []int
		target int
		expect []int
	}{
		{array: []int{1, 2, 3, 4, 5, 6, 7}, target: 3, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{array: []int{-1, -100, 3, 99}, target: 2, expect: []int{3, 99, -1, -100}},
		{array: []int{-1}, target: 2, expect: []int{-1}},
		{array: []int{-1}, target: 1, expect: []int{-1}},
		{array: []int{1, 2}, target: 3, expect: []int{2, 1}},
	}

	for _, test := range twoSumTest {
		rotate02(test.array, test.target)
		if !reflect.DeepEqual(test.array, test.expect) {
			t.Errorf("target:%v, expect:%v, get:%v", test.target, test.expect, test.array)
		}
	}
}
