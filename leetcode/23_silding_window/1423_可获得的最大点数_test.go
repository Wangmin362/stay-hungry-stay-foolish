package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/description/

// 题目分析：拿走k张扑克牌，剩余n-k张牌，由于拿走的和剩余的排总点数一定是常数，所以要想拿走的最大，那么剩余的保证最小即可
// 因此题目可以转化为求n-k张牌中，和最小的窗口

func maxScore(cardPoints []int, k int) int {
	total := 0
	for _, card := range cardPoints {
		total += card
	}
	cnt := len(cardPoints) - k // 求n-k窗口的最小值
	if cnt == 0 {
		return total
	}
	ans, sum := total, 0
	for idx, in := range cardPoints {
		sum += in

		if idx < cnt-1 {
			continue
		}

		ans = min(ans, sum)

		out := cardPoints[idx-cnt+1]
		sum -= out
	}
	return total - ans
}

func TestMaxScore(t *testing.T) {
	testdata := []struct {
		nums   []int
		k      int
		expect int
	}{
		//{nums: []int{1, 2, 3, 4, 5, 6, 1}, k: 3, expect: 12},
		{nums: []int{9, 7, 7, 9, 7, 7, 9}, k: 7, expect: 55},
	}

	for _, test := range testdata {
		get := maxScore(test.nums, test.k)
		if get != test.expect {
			t.Errorf("nums:%v, k:%v  expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
