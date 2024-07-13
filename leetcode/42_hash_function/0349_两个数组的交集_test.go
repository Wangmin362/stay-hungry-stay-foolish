package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/intersection-of-two-arrays/description/

// 解题思路：

func intersection(nums1 []int, nums2 []int) []int {
	mm := make(map[int]struct{})
	for idx := range nums1 {
		mm[nums1[idx]] = struct{}{}
	}
	var res []int
	for idx := range nums2 {
		if _, ok := mm[nums2[idx]]; ok {
			res = append(res, nums2[idx])
			delete(mm, nums2[idx])
		}
	}

	return res
}

func intersection01(nums1 []int, nums2 []int) []int {
	freCnt := func(nums []int) map[int]struct{} {
		cnt := make(map[int]struct{})
		for idx := range nums {
			cnt[nums[idx]] = struct{}{}
		}
		return cnt
	}

	res := map[int]struct{}{}
	nums1Cnt := freCnt(nums1)
	for idx := range nums2 {
		if _, ok := nums1Cnt[nums2[idx]]; ok {
			res[nums2[idx]] = struct{}{}
		}
	}
	var r []int
	for key := range res {
		r = append(r, key)
	}

	return r
}

func TestIntersection(t *testing.T) {
	var testdata = []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, expect: []int{2}},
	}

	for _, test := range testdata {
		get := intersection(test.nums1, test.nums2)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("num1:%v, num2:%v, expect:%v, get:%v", test.nums1, test.nums2, test.expect, get)
		}
	}

}
