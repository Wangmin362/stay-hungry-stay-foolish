package _1_array

import (
	"reflect"
	"testing"
)

// 题目：给你一个按非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 接替思路：考虑到负数的平方之后，其实是一个正数。平方之后，最大的数字一定是在左右两边，此时我们只需要把最大的数往新数组里面放即可

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums)) // 直接开辟等大空间，免得扩容浪费时间
	idx := len(nums) - 1

	left := 0
	right := len(nums) - 1
	for left <= right { // left=right时也是有效的
		ll := nums[left] * nums[left]
		rr := nums[right] * nums[right]
		if ll >= rr { // 取左边的数
			res[idx] = ll
			left++
		} else { // 取右边的数字
			res[idx] = rr
			right--
		}
		idx--
	}
	return res
}
func TestSortedSquares(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{9}, expect: []int{81}},
		{array: []int{0}, expect: []int{0}},
		{array: []int{-5, 0}, expect: []int{0, 25}},
		{array: []int{-5, -3, 0, 4, 6}, expect: []int{0, 9, 16, 25, 36}},
	}

	for _, test := range twoSumTest {
		get := sortedSquares(test.array)
		if !reflect.DeepEqual(test.expect, get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
