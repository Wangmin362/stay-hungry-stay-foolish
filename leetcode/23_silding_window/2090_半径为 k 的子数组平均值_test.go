package _1_array

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/k-radius-subarray-averages/description/

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	for idx := range nums { // 先直接全部初始化为-1
		res[idx] = -1
	}
	sum, left, idx, right := 0, 0, k, 2*k
	length := 2*k + 1
	if idx+k < len(nums) {
		for i := 0; i < length; i++ {
			sum += nums[i]
		}
	}

	for right < len(nums) {
		ave := sum / length
		res[idx] = ave
		idx++
		sum -= nums[left]
		left++
		right++
		if right < len(nums) {
			sum += nums[right]
		}

	}

	return res
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
