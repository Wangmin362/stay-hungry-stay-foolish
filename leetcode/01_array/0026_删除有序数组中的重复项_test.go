package _1_array

import (
	"testing"
)

// 解法一：直接遍历所有元素，然后每遍历到一个目标元素，直接把其后面的所有元素往前移动一格，时间复杂度为O(N^2)，空间复杂度为O(1)
// 解法二：直接遍历所有元素，然后每遍历到一个目标元素，直接最后位置的元素覆盖当前位置元素，时间复杂度为O(N)，空间复杂度为O(1)

func removeDuplicates(nums []int) int {
	return 0
}

func TestRemoveDuplicates(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{3, 2, 2, 3}, expect: []int{3, 2, 3}},
		{array: []int{0, 1, 2, 2, 3, 0, 4, 2}, expect: []int{0, 1, 2, 3, 0, 4, 2}},
	}

	for _, test := range twoSumTest {
		get := removeDuplicates(test.array)
		if get != len(test.expect) {
			t.Errorf("expect:%v, get:%v", test.expect, get)
		}

		for i := 0; i < get; i++ {
			if test.array[i] != test.expect[i] {
				t.Errorf("expect:%v, get:%v", test.expect, get)
			}
		}
	}
}
