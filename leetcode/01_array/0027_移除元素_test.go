package _1_array

import (
	"reflect"
	"testing"
)

// 解法一：直接遍历所有元素，然后每遍历到一个目标元素，直接把其后面的所有元素往前移动一格，时间复杂度为O(N^2)，空间复杂度为O(1)
// 解法二：直接遍历所有元素，然后每遍历到一个目标元素，直接最后位置的元素覆盖当前位置元素，时间复杂度为O(N)，空间复杂度为O(1)

func removeElement02(nums []int, val int) int {
	total := len(nums)
	for idx := 0; idx < total; idx++ {
		if nums[idx] == val {
			nums[idx] = nums[total-1]
			total--
			idx-- // 由于最后的元素有可能就是目标元素，因此需要索引减一，再次判断
		}
	}
	return total
}

func TestRemoveElement(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		expect int
	}{
		{array: []int{3, 2, 2, 3}, target: 3, expect: 2},
		{array: []int{0, 1, 2, 2, 3, 0, 4, 2}, target: 2, expect: 5},
	}

	for _, test := range twoSumTest {
		total := removeElement02(test.array, test.target)
		if !reflect.DeepEqual(total, test.expect) {
			t.Errorf("arr:%v, target:%v, expect:%v, get:%v", test.array, test.target, test.expect, total)
		}
	}
}
