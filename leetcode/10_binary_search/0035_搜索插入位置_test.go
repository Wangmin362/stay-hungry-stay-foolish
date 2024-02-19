package _0_binary_search

import (
	"testing"
)

// 分析：使用二分查找，每次看中间位置的元素是否比目标值大，如果比目标值大，那么继续从左边集合中搜索，如果比目标值小，那么从右边搜索
// 接替思路参考：https://leetcode.cn/problems/search-insert-position/solutions/1705568/by-carlsun-2-2dlr/
func searchInsert01(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	// 选择区间为[left, right]

	left := 0
	right := len(nums) - 1
	for left <= right { // 一直在[left, right]中找那个数，当left=right时，[left, right]依然有效
		middle := left + (right-left)>>1 // 防止数组越界，并且使用位运算加速
		if nums[middle] > target {
			right = middle - 1 // 中位数比目标值大，因此在左边区间找[left, middle-1]
		} else if nums[middle] < target { // 中位数比目标值小，因此在右边区间找[middle+1, right]
			left = middle + 1
		} else {
			return middle
		}
	}

	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  [0, -1]
	// 目标值等于数组中某一个元素  return middle;
	// 目标值插入数组中的位置 [left, right]，return  right + 1
	// 目标值在数组所有元素之后的情况 [left, right]，因为是右闭区间，所以 return right + 1
	return right + 1
}

func searchInsert02(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	// 选择区间为：[left, right)

	left := 0
	right := len(nums)
	for left < right { // 一直在[left, right)中找那个数，当left=right时，[left, right)无效，所以不能写等于
		middle := left + (right-left)>>1 // 防止数组越界，并且使用位运算加速
		if nums[middle] > target {
			right = middle // 中位数比目标值大，因此在左边区间找[left, middle)
		} else if nums[middle] < target { // 中位数比目标值小，因此在右边区间找[middle+1, right)
			left = middle + 1
		} else {
			return middle
		}
	}

	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  [0, 0]
	// 目标值等于数组中某一个元素  return middle;
	// 目标值插入数组中的位置 [left, right)，return  right
	// 目标值在数组所有元素之后的情况 [left, right)，因为是右闭区间，所以 return right
	return right
}

func TestSearch(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{array: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{array: []int{1, 3, 5, 6}, target: 1, expect: 0},
		{array: []int{1, 3, 5, 6}, target: 0, expect: 0},
		{array: []int{1, 3, 5, 6}, target: 7, expect: 4},
		{array: []int{1, 3, 5, 10}, target: 7, expect: 3},
	}

	for _, test := range twoSumTest {
		get := searchInsert02(test.array, test.target)
		if test.expect != get {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
