package _1_array

import (
	"testing"
)

// 解题思路：只需要从数组的最后一个元素开始遍历，如果元素不是9，直接加一返回；如果是9，置为0，继续往前遍历，直到找到不是9的元素，加一返回；
// 如果遍历完数组还没返回，说明是类似99这种情况，需要进位，直接在数组前面插入1即可

func plusOne1(digits []int) []int {
	idx := len(digits) - 1
	for idx >= 0 {
		if digits[idx] != 9 { // 如果最后一个元素不是9，直接加一返回
			digits[idx] += 1
			return digits
		} else { // 当前元素是9，置为0，继续往前遍历
			digits[idx] = 0
			idx--
		}
	}
	if idx < 0 {
		// 如果遍历完数组还没返回，说明是类似99这种情况，需要进位
		digits = append([]int{1}, digits...)
	}
	return digits
}

func TestPlusOne(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{9}, expect: []int{1, 0}},
		{array: []int{9, 9}, expect: []int{1, 0, 0}},
		{array: []int{3, 3, 3}, expect: []int{3, 3, 4}},
	}

	for _, test := range twoSumTest {
		get := plusOne1(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, test.array)
		}

		for i := 0; i < len(get); i++ {
			if get[i] != test.expect[i] {
				t.Fatalf("expect:%v, get:%v", test.expect, test.array)
			}
		}
	}
}
