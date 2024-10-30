package _1_array

import "sort"

// 解法一：直接合并两个数组，然后找中间的数字
func findMedianSortedArraysMerge(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}

	res := nums1[len(nums1)/2] + nums1[len(nums1)/2-1]
	return float64(res) / float64(2)
}
