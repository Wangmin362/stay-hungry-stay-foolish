package _1_array

import (
	"reflect"
	"sort"
	"testing"
)

// 题目：https://leetcode.cn/problems/3sum/description/

// TODO 此题还需要重新做

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for idx := 0; idx < len(nums); idx++ {
		left := idx + 1
		right := len(nums) - 1

		if nums[idx] > 0 { // 剪枝操作，因为第一个数都大于零了，后面的书又比这个数大，那么这三个数相加一定不为零
			return res
		}

		if idx > 0 && nums[idx] == nums[idx-1] { // 去重
			continue
		}

		for left < right {
			total := nums[idx] + nums[left] + nums[right]
			b, c := nums[left], nums[right]
			if total == 0 {
				res = append(res, []int{nums[idx], nums[left], nums[right]})
				for left < right && nums[left] == b {
					left++
				}
				for left < right && nums[right] == c {
					right--
				}
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
