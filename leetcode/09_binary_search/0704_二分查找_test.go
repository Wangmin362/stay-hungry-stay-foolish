package _9_binary_search

import (
	"testing"
)

// 目标：给定一个有序数组，没有重复的元素，在给定一个target，找到给定target在数组中的下标位置，如果不存在， 直接返回-1；
// 分析：由于是有序数组，且没有重复，所以可以直接使用二分查找，以log(n)的速度找到。
// 解题思路：解法一：区间定义为[left, right]全比去结案，解法二：[left, right),区间定义为左闭右开的区间。
// TODO 思考，如果存在重复元素，并且返回给定target的所有位置应该如何做？

// 先实现第一种，全闭区间 [left, right]， 那么中位数大于target，则right=middle-1, 如果中位数小于target，则left=middle+1
func searchAllClose(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { // 由于是全闭区间，所以[left, right]是有效的，该区间只包含了一个数字
		mid := left + (right-left)>>1 // 等价于 (right + left) / 2，这里为了防止溢出
		if nums[mid] > target {       // 中间的位置大于target，说明target在中间位置的左边，所以区间取中间位置左边
			right = mid - 1
		} else if nums[mid] < target { // 中间的位置小于target，说明target在中间位置右边，区间取中间位置的右边
			left = mid + 1
		} else { // 说明中间这个位置的数就是target
			return mid
		}
	}
	return -1 // 如果left > right，说明已经遍历完了，此时还没找到target，说明数组中不存在target，直接返回-1
}

// 这里实现第二种，即[left, right)左闭右开区间。mid为中间位置，若mid<target，说明target在右边，那么left=mid+1；若mid>target，说明target
// 在左边，那么right = mid
func searchRightOpen(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right { // 因为考虑的是[left, right)区间，因此left=right时，区间无效，因此不考虑相等
		mid := left + (right-left)>>1
		if nums[mid] > target { // 中间的数大于target，所以取左边的区间，由于是左开右闭，因此时mid，而不是mid-1
			right = mid
		} else if nums[mid] < target { // 中间的数小于target，所以取右边的区间，由于是左开右闭，因此时mid+1
			left = mid + 1
		} else {
			return mid
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
		get := searchRightOpen(test.array, test.target)
		if test.expect != get {
			t.Fatalf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
