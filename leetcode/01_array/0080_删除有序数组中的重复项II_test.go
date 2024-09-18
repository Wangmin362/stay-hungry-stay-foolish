package _1_array

import (
	"reflect"
	"testing"
)

func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	k := 1
	for i := 2; i < len(nums); i++ {
		// 当前元素和k位置以及k-1位置不同的话就保留，否则，直接移除
		if nums[i] != nums[k] || (nums[i] != nums[k-1]) {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func TestRemoveDuplicatesII(t *testing.T) {
	var testdata = []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 1, 1, 1, 2, 2, 2, 3}, want: []int{1, 1, 2, 2, 3}},
		{nums: []int{1, 1, 1, 2, 2, 3}, want: []int{1, 1, 2, 2, 3}},
	}
	for _, tt := range testdata {
		get := removeDuplicatesII(tt.nums)
		if !reflect.DeepEqual(tt.nums[:get], tt.want) {
			t.Fatalf("want:%v, get:%v", tt.want, tt.nums[:get])
		}
	}
}
