package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/

func maximumSubarraySum(nums []int, k int) int64 {
	cache := make(map[int]int)
	ans, sum := 0, 0
	for idx, in := range nums {
		sum += in
		cache[in]++

		if idx < k-1 {
			continue
		}

		if len(cache) == k { // 说明各不相同的元素
			ans = max(ans, sum)
		}

		out := nums[idx-k+1]
		sum -= out
		if cnt := cache[out]; cnt == 1 {
			delete(cache, out)
		} else {
			cache[out]--
		}
	}
	return int64(ans)
}

func TestMaximumSubarraySum(t *testing.T) {
	testdata := []struct {
		nums   []int
		k      int
		expect int64
	}{
		{nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3, expect: 15},
	}

	for _, test := range testdata {
		get := maximumSubarraySum(test.nums, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
