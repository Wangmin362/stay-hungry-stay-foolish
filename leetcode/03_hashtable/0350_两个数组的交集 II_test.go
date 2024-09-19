package _0_basic

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/intersection-of-two-arrays/description/

// 解题思路：

func intersectii(nums1 []int, nums2 []int) []int {
	mm := make(map[int]int)
	for idx := range nums1 {
		mm[nums1[idx]] += 1
	}

	var res []int
	for idx := range nums2 {
		cnt, ok := mm[nums2[idx]]
		if ok && cnt > 0 {
			res = append(res, nums2[idx])
			mm[nums2[idx]] -= 1
		}
	}

	return res
}

func TestIntersectii(t *testing.T) {
	var testdata = []struct {
		nums1  []int
		nums2  []int
		expect []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, expect: []int{2, 2}},
	}

	for _, test := range testdata {
		get := intersectii(test.nums1, test.nums2)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("num1:%v, num2:%v, expect:%v, get:%v", test.nums1, test.nums2, test.expect, get)
		}
	}

}
