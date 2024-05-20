package _1_array

import (
	"reflect"
	"sort"
	"testing"
)

// 题目：https://leetcode.cn/problems/3sum/description/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for idx := 0; idx < len(nums); idx++ {
		left := idx + 1
		right := len(nums) - 1

		for left < right {
			total := nums[idx] + nums[left] + nums[right]
			if total == 0 {
				res = append(res, []int{nums[idx], nums[left], nums[right]})
				left++
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}

func TestThreeSum(t *testing.T) {
	var teatdata = []struct {
		array  []int
		expect [][]int
	}{
		{array: []int{-1, 0, 1, 2, -1, -4}, expect: [][]int{{0, 1}}},
	}

	for _, test := range teatdata {
		sum01 := threeSum(test.array)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("arr:%v, expect:%v, get:%v", test.array, test.expect, sum01)
		}
	}
}
