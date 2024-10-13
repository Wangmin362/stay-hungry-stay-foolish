package _0_basic

import (
	"testing"
)

func largestPerimeter(nums []int) int {
	return 0
}

func TestLargestPerimeter(t *testing.T) {
	var testsdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3}, want: 4},
	}
	for _, tt := range testsdata {
		get := largestPerimeter(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
