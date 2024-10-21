package _9_binary_search

import "testing"

// 先通过二分搜索搜索数组中的最小值，然后在两边查找target
func search(nums []int, target int) int {
	findMin := func(nums []int) int { // 最小值一定存在
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] < nums[right] {
				right = mid // 最小值要么在左侧，要么这里就是最小值
			} else {
				left = mid + 1 // 说明最小值一定在右侧，当前点不可能是最小值
			}
		}
		return right
	}

	binarySearch := func(nums []int, start, end, target int) int {
		left, right := start, end
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	minIndex := findMin(nums) // 找到最小值索引所在位置
	idx := binarySearch(nums, minIndex, len(nums)-1, target)
	if idx != -1 {
		return idx
	}

	return binarySearch(nums, 0, minIndex-1, target)
}

func TestSearch0033(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		target int
		want   int
	}{
		{array: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
	}

	for _, test := range twoSumTest {
		get := search(test.array, test.target)
		if get != test.want {
			t.Fatalf("array:%v, target:%v, want:%v, get:%v", test.array, test.target, test.want, get)
		}
	}
}
