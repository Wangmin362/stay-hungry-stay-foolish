package _0_basic

import (
	"reflect"
	"sort"
	"testing"
)

func minSubsequence(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var sum int
	for _, num := range nums {
		sum += num
	}

	sort.Ints(nums)
	mid := sum >> 1

	left, right, seqSum := len(nums)-1, len(nums)-1, 0
	for ; left >= 0; left-- {
		seqSum += nums[left]
		if seqSum > mid {
			break
		}
	}

	l, r := left, right
	for left < right { // 逆序输出
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	return nums[l : r+1]
}

func TestMinSubsequence(t *testing.T) {
	var testsdata = []struct {
		nums []int
		want []int
	}{
		//{nums: []int{4, 3, 10, 9, 8}, want: []int{10, 9}},
		{nums: []int{1, 7, 4, 7, 1, 9, 4, 8, 8}, want: []int{9, 8, 8}},
	}
	for _, tt := range testsdata {
		get := minSubsequence(tt.nums)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("nums:%v,  want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
