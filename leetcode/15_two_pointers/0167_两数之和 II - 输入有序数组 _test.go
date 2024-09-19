package _0_basic

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/?envType=study-plan-v2&envId=top-interview-150

// 使用对撞指针
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{0, 0}
}

func TestTwoSum(t *testing.T) {
	var testData = []struct {
		numbers []int
		target  int
		want    []int
	}{
		{numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
	}

	for _, test := range testData {
		get := twoSum(test.numbers, test.target)
		if !reflect.DeepEqual(get, test.want) {
			t.Errorf("numbers:%v, target:%v, want:%v, get:%v", test.numbers, test.target, test.want, get)
		}
	}

}
