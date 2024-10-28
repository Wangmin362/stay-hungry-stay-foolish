package _0_basic

import (
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res int
	diff := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			} else if sum > target {
				k--
				if diff > sum-target {
					diff = sum - target
					res = sum
				}
			} else {
				j++
				if diff > target-sum {
					diff = target - sum
					res = sum
				}
			}
		}
	}
	return res
}

func TestThreeSumClosest(t *testing.T) {
	var teatdata = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, target: -2, want: -2},
	}

	for _, test := range teatdata {
		get := threeSumClosest(test.nums, test.target)
		if get != test.want {
			t.Errorf("nums:%v, target:%v, want:%v,  get:%v", test.nums, test.target, test.want, get)
		}
	}
}
