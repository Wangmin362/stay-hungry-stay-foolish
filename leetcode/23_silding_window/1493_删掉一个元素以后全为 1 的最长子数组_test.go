package _1_array

import "testing"

func longestSubarray(nums []int) int {
	res, one, zero := 0, 0, 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		} else {
			one++
		}
		for zero > 1 {
			if nums[left] == 0 {
				zero--
			} else {
				one--
			}
			left++
		}
		if zero == 1 {
			res = max(res, one) // 把零删除即可
		} else {
			res = max(res, one-1) // 全是1，但是必须删除一个数字
		}

	}
	return res
}

func TestLongestSubarray(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 1, 0, 0, 1}, want: 3},
		{nums: []int{0, 0, 0, 0}, want: 0},
	}
	for _, tt := range testdata {
		get := longestSubarray(tt.nums)
		if get != tt.want {
			t.Errorf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}

}
