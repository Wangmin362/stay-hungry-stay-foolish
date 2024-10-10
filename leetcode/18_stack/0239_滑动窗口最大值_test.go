package _1_array

import (
	"math"
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/sliding-window-maximum/description/

func maxSlidingWindow(nums []int, k int) []int {
	getMax := func(nums []int, start, end int) int {
		iMax := math.MinInt
		for idx := start; idx <= end; idx++ {
			if nums[idx] > iMax {
				iMax = nums[idx]
			}
		}
		return iMax
	}
	var res []int
	for idx := 0; idx <= len(nums)-k; idx++ {
		res = append(res, getMax(nums, idx, idx+k-1))
	}
	return res
}

type queue []int

func (r *queue) pop(x int) {
	if len(*r) > 0 && x == r.front() {
		*r = (*r)[1:]
	}
}
func (r *queue) push(x int) {
	for len(*r) > 0 && (*r)[len(*r)-1] < x {
		*r = (*r)[:len(*r)-1]
	}
	*r = append(*r, x)
}
func (r queue) front() int {
	if len(r) > 0 {
		return r[0]
	}
	return -1
}

func maxSlidingWindow02(nums []int, k int) []int {
	var res []int
	qe := queue{}
	for i := 0; i < k; i++ {
		qe.push(nums[i])
	}
	res = append(res, qe.front())

	for i := 0; i+k < len(nums); i++ {
		qe.pop(nums[i])
		qe.push(nums[i+k])
		res = append(res, qe.front())
	}
	return res
}

func TestMaxSlidingWindow(t *testing.T) {
	var teatdata = []struct {
		nums   []int
		k      int
		expect []int
	}{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, expect: []int{3, 3, 5, 5, 6, 7}},
	}

	for _, test := range teatdata {
		get := maxSlidingWindow02(test.nums, test.k)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("nums: %v, k: %v, expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
