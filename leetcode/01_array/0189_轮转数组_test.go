package _1_array

import (
	"reflect"
	"testing"
)

// 对于go来说相当简单，直接两次append直接搞定，但是由于切片是不可变的，因此地址发生了改变。所以我们需要在不修改地址的情况下实现
func rotate01(nums []int, k int) {
	if k <= 0 {
		return
	}

	var res []int
	for cnt := 0; cnt < k; cnt++ {
		res = append(res, nums[len(nums)-1])
		res = append(res, nums[:len(nums)-1]...)
		nums = res
		res = make([]int, 0)
	}
}

// 部分case没有通过
func rotate02(nums []int, k int) {
	if k <= 0 {
		return
	}

	if len(nums) < k {
		k = len(nums)
	}

	tmpArr := make([]int, k)
	for i := 0; i < k; i++ {
		tmpArr[i] = nums[len(nums)-k+i]
	}

	for i := 0; i < len(nums)-k; i++ { // 需要移动Len(nums) - k个元素
		nums[len(nums)-i-1] = nums[len(nums)-k-i-1]
	}

	for i := 0; i < k; i++ {
		nums[i] = tmpArr[i]
	}
}

func TestRotate(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect []int
	}{
		{array: []int{1, 2, 3, 4, 5, 6, 7}, target: 3, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{array: []int{-1, -100, 3, 99}, target: 2, expect: []int{3, 99, -1, -100}},
		{array: []int{-1}, target: 2, expect: []int{-1}},
		{array: []int{-1}, target: 1, expect: []int{-1}},
		{array: []int{1, 2}, target: 3, expect: []int{2, 1}},
	}

	for _, test := range twoSumTest {
		rotate02(test.array, test.target)
		if !reflect.DeepEqual(test.array, test.expect) {
			t.Errorf("target:%v, expect:%v, get:%v", test.target, test.expect, test.array)
		}
	}
}
