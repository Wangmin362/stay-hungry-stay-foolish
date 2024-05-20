package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

// 解题思路：从后往前合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Idx := m - 1
	nums2Idx := len(nums2) - 1
	totalIdx := n + m - 1
	for nums1Idx >= 0 && nums2Idx >= 0 {
		if nums1[nums1Idx] > nums2[nums2Idx] {
			nums1[totalIdx] = nums1[nums1Idx]
			nums1Idx--
		} else {
			nums1[totalIdx] = nums2[nums2Idx]
			nums2Idx--
		}
		totalIdx--
	}
	for nums2Idx >= 0 {
		nums1[totalIdx] = nums2[nums2Idx]
		totalIdx--
		nums2Idx--
	}
}

func TestMerge(t *testing.T) {
	var testdata = []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		//{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expect: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{4, 5, 6, 0, 0, 0}, m: 3, nums2: []int{1, 2, 3}, n: 3, expect: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range testdata {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.expect) {
			t.Errorf("nums1:%v, m:%v, nums2:%v, n:%v, expect:%v, get:%v", test.nums1, test.m, test.nums2, test.n, test.expect, test.nums1)
		}
	}
}
