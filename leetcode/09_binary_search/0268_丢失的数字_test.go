package _9_binary_search

import "testing"

// 解法一：O(N^2)的时间复杂度 + O(1)的空间复杂度，遍历扫描
func missingNumber01(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	total := len(nums)
	for i := 0; i <= total; i++ {
		hasNum := false
		for idx := range nums {
			if nums[idx] == i {
				hasNum = true
				break
			}
		}
		if !hasNum {
			return i
		}
	}
	return -1
}

// 解法二：Time O(N), Store O(N)  利用map
func missingNumber02(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	cache := make(map[int]struct{})

	for idx := range nums {
		cache[nums[idx]] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, exist := cache[i]; !exist {
			return i
		}
	}
	return -1
}

// TODO 使用异或运算 这TMD时人能想出来的？？？？？
func missingNumber04(nums []int) int {
	n := len(nums)
	ans := 0
	for i, num := range nums {
		// 0 ^ 1 ^... ^(n-1) ^ num(s)
		ans ^= i ^ num
	}
	return ans ^ n
}

func TestMissingNumber(t *testing.T) {
	var twoSumTest = []struct {
		array  []int
		expect int
	}{
		{array: []int{3, 0, 1}, expect: 2},
		{array: []int{0, 1}, expect: 2},
		{array: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, expect: 8},
		{array: []int{0}, expect: 1},
		{array: []int{}, expect: -1},
		{array: nil, expect: -1},
	}

	for _, test := range twoSumTest {
		get := missingNumber02(test.array)
		if test.expect != get {
			t.Errorf("arr:%v,  expect:%v, get:%v", test.array, test.expect, get)
		}
	}
}
