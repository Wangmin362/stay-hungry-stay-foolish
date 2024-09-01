package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/

func findMaxAverage(nums []int, k int) float64 {
	ans := float64(-99999999999)
	sum := 0
	for idx, in := range nums {
		// 处理有边界，进来直接加一
		sum += in

		// 判断初始条件，判断窗口是否到达约定窗口大小
		if idx < k-1 { // 题目保证K >= len(nums)
			continue
		}

		// 计算平均值，并更新
		avg := float64(sum) / float64(k)
		ans = max(ans, avg)

		// 处理窗口左边界
		out := nums[idx-k+1]
		sum -= out
	}
	return ans
}

func TestFindMaxAverage(t *testing.T) {
	testdata := []struct {
		nums   []int
		k      int
		expect float64
	}{
		//{nums: []int{1, 12, -5, -6, 50, 3}, k: 4, expect: 12.75},
		{nums: []int{4433, -7832, -5068, 4009, 2830, 6544, -6119, -7126, -780, -4254, -8249, -9168, 9492, 402, 5789, 6808, 8953, 5810, -7353, 7933, 4766, 5182, -3230, -1989, 5786, 6922, -4646, 4415, -9906, 807, -6373, 3370, 2604, 8751, -9173, -2668, -6876, 9500, 3465, -1900, 4134, -1758, -1453, -5201, -9825, 4469, -1999, -1108, 1836, 3923, 6796, -5252, 9863, -5997, -3251, 9596, -3404, -540, 2826, -1737, 3341, -3623, -9885, 2603, -5782, 8174, 2710, 6504, -4128}, k: 59, expect: 526.37288},
	}

	for _, test := range testdata {
		get := findMaxAverage(test.nums, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
