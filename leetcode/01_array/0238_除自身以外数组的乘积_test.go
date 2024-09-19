package _1_array

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150

// 解法一：暴力计算
// 解法二：前后缀分解

func productExceptSelf01(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		mul := 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			mul *= nums[j]
		}
		res[i] = mul
	}
	return res
}

// 定义pre[i]为nums[0:i-1]之所所有数字的乘积， 那么 pre[i] = pre[i-1] * nums[i-1]
// 定义suf[i]为nums[i+1:]之间所有数字的乘积，那么suf[i] = suf[i+1] * nums[i+1]
// 由上面推导：ans[i] = pre[i] * suf[i]
func productExceptSelf02(nums []int) []int {
	pre, suf := make([]int, len(nums)), make([]int, len(nums))
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	suf[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = pre[i] * suf[i]
	}

	return ans
}

func TestProductExceptSelf(t *testing.T) {
	var testdata = []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range testdata {
		get := productExceptSelf02(tt.nums)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("citations:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
