package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/description/

func numOfSubarrays0901(arr []int, k int, threshold int) int {
	ans, sum := 0, 0
	for idx, in := range arr {
		sum += in

		if idx < k-1 { // 窗口还不够
			continue
		}

		avg := sum / k
		if avg >= threshold {
			ans++
		}

		out := arr[idx-k+1]
		sum -= out
	}
	return ans
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
		get := numOfSubarrays0901(test.nums, test.k, test.threshold)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
