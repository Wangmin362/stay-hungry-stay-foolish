package _0_basic

import (
	"testing"
)

// 贪心思想：第一个子数组的代价已经确定，其实就是nums[0]，如果想要是代价最小，那么需要在后面的数组当中找到最小的两个数字即可
func minimumCost(nums []int) int {
	fir, sec := nums[1], nums[2]
	if fir > sec {
		fir, sec = sec, fir
	}
	for i := 3; i < len(nums); i++ { // 从第一个数字开始找
		if nums[i] < fir {
			sec = fir
			fir = nums[i]
		} else if nums[i] < sec {
			sec = nums[i]
		}
	}
	return nums[0] + fir + sec
}

func TestMinimumCost(t *testing.T) {
	var testsdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 12}, want: 6},
	}
	for _, tt := range testsdata {
		get := minimumCost(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
