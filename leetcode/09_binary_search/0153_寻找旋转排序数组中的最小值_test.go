package _9_binary_search

import "testing"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] { // 如果当前值小于最右边的元素，说明最小值一定在左边，所有右区间
			right = mid // 之所以为mid，是因为mid可能是最小值
		} else { // 说明右侧一定有拐点，或者说拐点一定在当前值的右侧，并且mid一定不是最小值，所以是Mid+1
			left = mid + 1
		}
	}
	return nums[right]
}

func TestFindMin(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{3, 4, 5, 1, 2}, want: 1},
	}

	for _, tt := range testdata {
		get := findMin(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
