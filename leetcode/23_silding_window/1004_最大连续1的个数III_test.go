package _1_array

import "testing"

// 解题思路参考：https://leetcode.cn/problems/max-consecutive-ones-iii/solutions/2126631/hua-dong-chuang-kou-yi-ge-shi-pin-jiang-yowmi/?envType=company&envId=huawei&favoriteSlug=huawei-all

// 解题思路：滑动窗口，求窗口中的0的个数小于等于k的情况下，窗口的最大值
func longestOnes(nums []int, k int) int {
	left, right, zero, res := 0, 0, 0, 0
	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
			for zero > k { // 说明窗口中的0大于规定的数量，此时需要移动左边窗口
				if nums[left] == 0 {
					zero--
				}
				left++
			}
		}
		res = max(right-left+1, res)
	}

	return res
}

func TestLongestOnes(t *testing.T) {
	var testdata = []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2, want: 6},
		{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3, want: 10},
	}
	for _, tt := range testdata {
		get := longestOnes(tt.nums, tt.k)
		if get != tt.want {
			t.Errorf("nums:%v, k:%v, want:%v, get:%v", tt.nums, tt.k, tt.want, get)
		}
	}
}
