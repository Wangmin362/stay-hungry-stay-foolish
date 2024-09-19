package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-interview-150

// 使用对撞指针
// 若左边的高度必有右边的短，那么一定是移动左边板子，因为移动右边容积一定更小
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, mul := 0, 0
	for left < right {
		if height[left] < height[right] {
			mul = (right - left) * height[left]
			left++
		} else {
			mul = (right - left) * height[right]
			right--
		}
		res = max(res, mul)
	}
	return res
}

func TestMaxArea(t *testing.T) {
	var testData = []struct {
		height []int
		want   int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
	}

	for _, tt := range testData {
		get := maxArea(tt.height)
		if get != tt.want {
			t.Errorf("height:%v, want:%v, get:%v", tt.height, tt.want, get)
		}
	}

}
