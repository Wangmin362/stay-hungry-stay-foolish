package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/

func numOfSubarrays(arr []int, k int, threshold int) int {
	if len(arr) <= 0 {
		return 0
	}
	sum := 0
	if len(arr) < k {
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		ave := float64(sum) / float64(len(arr))
		if ave >= float64(threshold) {
			return 1
		}
		return 0
	}
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	var res int
	l, r := 0, k-1
	for r < len(arr) {
		ave := float64(sum) / float64(k)
		if ave >= float64(threshold) {
			res++
		}

		sum -= arr[l]
		l++
		r++
		if r < len(arr) {
			sum += arr[r]
		}
	}

	return res
}

func TestNumOfSubarrays(t *testing.T) {
	testdata := []struct {
		nums      []int
		k         int
		threshold int
		expect    int
	}{
		{nums: []int{2, 2, 2, 2, 5, 5, 5, 8}, k: 3, threshold: 4, expect: 3},
	}

	for _, test := range testdata {
		get := numOfSubarrays(test.nums, test.k, test.threshold)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
