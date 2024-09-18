package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

// 解题思路：从后往前合并
func merge0918(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[idx] = nums1[m]
			m--
		} else {
			nums1[idx] = nums2[n]
			n--
		}
		idx--
	}
	for n >= 0 {
		nums1[idx] = nums2[n]
		idx--
		n--
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
		merge0918(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.expect) {
			t.Errorf("nums1:%v, m:%v, nums2:%v, n:%v, expect:%v, get:%v", test.nums1, test.m, test.nums2, test.n, test.expect, test.nums1)
		}
	}
}
