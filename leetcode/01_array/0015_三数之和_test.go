package _1_array

import (
	"reflect"
	"slices"
	"testing"
)

// 题目：https://leetcode.cn/problems/3sum/description/

func threeSum02(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	slices.Sort(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			return res
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
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
		{array: []int{-1, 0, 1, 2, -1, -4}, expect: [][]int{{0, 1, -1}, {-1, -1, 2}}},
	}

	for _, test := range teatdata {
		sum01 := threeSum02(test.array)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("arr:%v, expect:%v, get:%v", test.array, test.expect, sum01)
		}
	}
}
