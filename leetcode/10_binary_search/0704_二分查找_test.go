package _0_binary_search

import "testing"

// 分析：区间为[left, right]
func search(nums []int, target int) int {
	// 下面的两个检测条件可以不需要，因为根本就不满足for循环条件
	//if nums == nil {
	//	return -1
	//}
	//if len(nums) == 0 {
	//	return -1
	//}

	left := 0
	right := len(nums) - 1
	for left <= right { // 由于是闭区间[left, right]，因此left=right时，区间依然有效，即只有一个值
		middle := left + (right-left)>>1
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1 // 由于中位数比目标大，因此目标一定在中位数左边，所以区间为[left, middle-1]
		} else if nums[middle] < target {
			left = middle + 1 // 由于中位数比目标小，因此目标一定在中位数右边，所以区间为[middle+1, left]
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{-1, 0, 3, 5, 9, 12}, target: -1, expect: 0},
		{array: []int{-1, 0, 3, 5, 9, 12}, target: 3, expect: 2},
		{array: []int{-1, 0, 3, 5, 9, 12}, target: 12, expect: 5},
		{array: []int{-1, 0, 3, 5, 9, 12}, target: -9, expect: -1},
		{array: []int{-1, 0, 3, 5, 9, 12}, target: 20, expect: -1},
		{array: nil, target: 20, expect: -1},
		{array: []int{}, target: 20, expect: -1},
	}

	for _, test := range twoSumTest {
		get := search(test.array, test.target)
		if test.expect != get {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
