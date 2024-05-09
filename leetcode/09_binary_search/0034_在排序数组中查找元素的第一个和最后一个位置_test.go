package _9_binary_search

import (
	"testing"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////下面采用左闭右闭区间解决此问题//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 首先采用闭区间解决这个问题，即循环不变量为[left, right]， mid=(right+left)/2
// 若nums[mid] > target，说明target在中位数左边，此时right=mid-1
// 若nums[mid] < target, 说明target在中位数右边，此时left=mid+1
// 若nums[mid] = target, 说明找到了target，此时很有可能target有多个，此时需要继续向左右两边查找
// 左边只需要mid--，直到找到那个不等于target的位置即可，后一个位置就是第一个下标
// 右边只需要mid++，直到找到那个不等于target的位置即可，前一个位置就是最后一个下标

func searchRangeAllClose(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 说明target在中位数的左边，取左边的区间
			right = mid - 1
		} else if nums[mid] < target { // 说明target在中位数的右边，去右边区间
			left = mid + 1
		} else { // 说明相等，此时左右两边很有可能还有相同的数字，左边递减，右边递加查找即可
			// 查找左边边界
			leftIdx := mid - 1
			leftTarget := mid
			for leftIdx >= 0 {
				if nums[leftIdx] != target {
					leftTarget = leftIdx + 1 // 一定是后面的一个位置
					break
				}
				leftIdx--
			}
			if leftIdx < 0 {
				leftTarget = 0
			}

			rightIdx := mid + 1
			rightTarget := mid
			for rightIdx < len(nums) {
				if nums[rightIdx] != target {
					rightTarget = rightIdx - 1
					break
				}
				rightIdx++
			}
			if rightIdx >= len(nums) {
				rightTarget = len(nums) - 1
			}
			return []int{leftTarget, rightTarget}
		}
	}

	return []int{-1, -1}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////下面采用左闭右开区间解决此问题//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func searchRangeRightOpen(nums []int, target int) []int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 说明target在中位数的左边，取左边的区间
			right = mid
		} else if nums[mid] < target { // 说明target在中位数的右边，去右边区间
			left = mid + 1
		} else { // 说明相等，此时左右两边很有可能还有相同的数字，左边递减，右边递加查找即可
			// 查找左边边界
			leftIdx := mid - 1
			leftTarget := mid
			for leftIdx >= 0 {
				if nums[leftIdx] != target {
					leftTarget = leftIdx + 1 // 一定是后面的一个位置
					break
				}
				leftIdx--
			}
			if leftIdx < 0 {
				leftTarget = 0
			}

			rightIdx := mid + 1
			rightTarget := mid
			for rightIdx < len(nums) {
				if nums[rightIdx] != target {
					rightTarget = rightIdx - 1
					break
				}
				rightIdx++
			}
			if rightIdx >= len(nums) {
				rightTarget = len(nums) - 1
			}
			return []int{leftTarget, rightTarget}
		}
	}

	return []int{-1, -1}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////左闭右闭 二分加速//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 这里的查找思路非常简单，其实就是先通过二分找到一个等于target的数，然后这个数的左右两次很有可能还有target，因此还需要使用二分查找

func searchRangeAllCloseSpeed(nums []int, target int) []int {
	leftIdx := -1
	rightIdx := -1

	// 查找左边界
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 说明target在中位数的左边，取左边的区间
			right = mid - 1
		} else if nums[mid] < target { // 说明target在中位数的右边，去右边区间
			left = mid + 1
		} else { // 说明相等，此时左右两边很有可能还有相同的数字，左边递减，右边递加查找即可
			leftIdx = mid
			right = mid - 1 // 继续使用二分法查找左边等于target的索引，直到退出循环
		}
	}

	// 查找左边界
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target { // 说明target在中位数的左边，取左边的区间
			right = mid - 1
		} else if nums[mid] < target { // 说明target在中位数的右边，去右边区间
			left = mid + 1
		} else { // 说明相等，此时左右两边很有可能还有相同的数字，左边递减，右边递加查找即可
			rightIdx = mid
			left = mid + 1 // 继续使用二分法查找左边等于target的索引，直到退出循环
		}
	}

	return []int{leftIdx, rightIdx}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////左闭右开 二分加速//////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestSearchRange(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect []int
	}{
		{array: []int{2, 2}, target: 2, expect: []int{0, 1}},
		{array: []int{2}, target: 2, expect: []int{0, 0}},
		{array: []int{8, 8, 8, 8, 8, 10}, target: 8, expect: []int{0, 4}},
		{array: []int{5, 7, 7, 8, 8, 10}, target: 8, expect: []int{3, 4}},
		{array: []int{5, 7, 7, 8, 8, 8, 10}, target: 8, expect: []int{3, 5}},
	}

	for _, test := range twoSumTest {
		get := searchRangeAllCloseSpeed(test.array, test.target)
		if len(get) != 2 || get[0] != test.expect[0] || get[1] != test.expect[1] {
			t.Fatalf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, get)
		}
	}
}
