package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/description/

func maxSum(nums []int, m int, k int) int64 {
	cache := make(map[int]int, k) // 窗口中最多k个元素
	ans, sum := 0, 0
	for idx, in := range nums {
		sum += in
		cache[in]++

		if idx < k-1 {
			continue
		}
		if len(cache) >= m {
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

func TestMaxSum(t *testing.T) {
	testdata := []struct {
		nums   []int
		m      int
		k      int
		expect int64
	}{
		{nums: []int{2, 6, 7, 3, 1, 7}, m: 3, k: 4, expect: 18},
	}

	for _, test := range testdata {
		get := maxSum(test.nums, test.m, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
