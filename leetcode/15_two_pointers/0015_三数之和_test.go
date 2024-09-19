package _0_basic

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/3sum/description/

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

// 暴力搜索
func threeSum01(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}

// 先排序，在使用碰撞指针
func threeSum03(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 { // 后面的数字一定都是大于零的，相加之和一定大于零，不可能是负数，直接跳过
			break
		}

		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > target {
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
		{array: []int{0, 0, 0, 0}, expect: [][]int{{0, 0, 0}}},
	}

	for _, test := range teatdata {
		sum01 := threeSum03(test.array)
		if !reflect.DeepEqual(sum01, test.expect) {
			t.Errorf("arr:%v, expect:%v, get:%v", test.array, test.expect, sum01)
		}
	}
}
