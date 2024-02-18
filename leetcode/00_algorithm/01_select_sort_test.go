package _0_algorithm

import (
	"reflect"
	"testing"
)

func getMin(arr []int) (int, int) {
	if arr == nil {
		return -1, -1
	}
	if len(arr) == 0 {
		return -1, -1
	}

	min := arr[0]
	minIdx := 0
	for idx := range arr {
		if idx == 0 {
			continue
		}
		if arr[idx] < min {
			min = arr[idx]
			minIdx = idx
		}
	}
	return min, minIdx
}

// 选择排序实现，没有在原地完成，并且使用了额外的空间, 时间复杂度未O(N^2)
func SelectSort01(arr []int) []int {
	if arr == nil {
		return nil
	}

	tmp := make([]int, len(arr))

	total := len(arr)
	for idx := 0; idx < total; idx++ {
		min, minIdx := getMin(arr)
		tmp[idx] = min
		tmpArr := make([]int, 0, len(arr))
		tmpArr = append(tmpArr, arr[:minIdx]...)
		tmpArr = append(tmpArr, arr[minIdx+1:]...)
		arr = tmpArr
	}

	return tmp
}

// 时间复杂度为O(N^2)，空间复杂度为O(1)
func SelectSort02(arr []int) []int {
	if arr == nil {
		return nil
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}

func TestSort(t *testing.T) {
	var sortTest = []struct {
		array  []int
		expect []int
	}{
		{array: []int{2, 7, 11, 15}, expect: []int{2, 7, 11, 15}},
		{array: []int{7, 2, 11, 15}, expect: []int{2, 7, 11, 15}},
		{array: []int{11, 7, 2, 15}, expect: []int{2, 7, 11, 15}},
		{array: []int{7, 11, 15, 2}, expect: []int{2, 7, 11, 15}},
		{array: []int{}, expect: []int{}},
		{array: []int{1}, expect: []int{1}},
		{array: nil, expect: nil},
	}

	for _, test := range sortTest {
		get := SelectSort02(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("arr:%v, expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
