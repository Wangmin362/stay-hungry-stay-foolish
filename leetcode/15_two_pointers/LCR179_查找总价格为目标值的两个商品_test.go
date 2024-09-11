package _0_basic

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/description/?envType=problem-list-v2&envId=two-pointers&difficulty=EASY

func twoSumLCR179(price []int, target int) []int {
	left, right := 0, len(price)-1
	for left < right {
		sum := price[left] + price[right]
		if sum == target {
			return []int{price[left], price[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

func TestTwoSumLCR(t *testing.T) {
	var testData = []struct {
		price  []int
		target int
		want   []int
	}{
		{price: []int{3, 9, 12, 15}, target: 18, want: []int{3, 15}},
	}

	for _, tt := range testData {
		get := twoSumLCR179(tt.price, tt.target)
		if !reflect.DeepEqual(get, tt.want) {
			t.Fatalf("price:%v, target:%v, want:%v, get:%v", tt.price, tt.target, tt.want, get)
		}
	}

}
