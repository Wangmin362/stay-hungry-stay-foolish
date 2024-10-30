package _1_array

import "sort"

// 解法一：直接排序后比较
func findUnsortedSubarraySort(nums []int) int {
	bak := make([]int, len(nums))
	copy(bak, nums)
	sort.Ints(bak)

	left, right := 0, len(nums)-1
	for left <= right && bak[left] == nums[left] {
		left++
	}

	for left <= right && bak[right] == nums[right] {
		right--
	}

	if right-left < 1 {
		return 0
	}

	return right - left + 1
}
