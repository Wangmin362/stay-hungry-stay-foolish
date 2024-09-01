package _1_array

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/k-radius-subarray-averages/description/

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		ans[idx] = -1
	}

	sum := 0
	for idx, in := range nums {
		sum += in

		if idx < 2*k {
			continue
		}

		avg := sum / (2*k + 1)
		ans[idx-k] = avg

		out := nums[idx-2*k]
		sum -= out
	}
	return ans
}

func TestGetAverages(t *testing.T) {
	testdata := []struct {
		num    []int
		k      int
		expect []int
	}{
		{num: []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, k: 3, expect: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
		{num: []int{100000}, k: 0, expect: []int{100000}},
		{num: []int{8}, k: 800, expect: []int{-1}},
	}

	for _, test := range testdata {
		get := getAverages(test.num, test.k)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("num:%v, t:%v  expect:%v, get:%v", test.num, test.k, test.expect, get)
		}
	}
}
