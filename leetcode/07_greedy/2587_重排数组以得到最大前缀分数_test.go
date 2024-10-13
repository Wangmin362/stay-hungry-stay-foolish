package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：按照从小到大排序，然后从右边计算前缀和
func maxScore(nums []int) int {
	sort.Ints(nums)
	res, prefixSum := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		if prefixSum+nums[i] <= 0 {
			break
		}
		prefixSum += nums[i]
		res++
	}

	return res
}
func TestMaxScore(t *testing.T) {
	var testsdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{2, -1, 0, 1, -3, 3, -3}, want: 6},
	}
	for _, tt := range testsdata {
		get := maxScore(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
