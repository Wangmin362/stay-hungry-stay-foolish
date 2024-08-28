package _9_binary_search

import "slices"

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/description/

// 题目分析：其实就是要求arr1的每个元素和arr2每个元素相减的绝对值满足举例即可

// 思路一：两个for循环，统计，时间复杂度为O(m * n)
// 思路二：第二个数组排序，然后使用O(m * log(n))的实践复杂度

// 思路一解题
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	absDiff := func(x, y int) int {
		if x > y {
			return x - y
		} else {
			return y - x
		}
	}

	var cnt int
	for _, a1 := range arr1 {
		valid := true
		for _, a2 := range arr2 {
			if absDiff(a1, a2) <= d {
				valid = false
				break
			}
		}
		if valid {
			cnt++
		}
	}
	return cnt
}

// 思路二解题：
func findTheDistanceValue02(arr1 []int, arr2 []int, d int) int {
	// 先排序
	slices.Sort(arr2)
	leftBoarder := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}

	var cnt int
	for _, a1 := range arr1 {
		min, max := a1-d-1, a1+d+1 // 若arr2数组中存在[min, max]之间的数字，说明肯定不符合距离值
		lIdx := leftBoarder(arr2, min+1) - 1
		rIdx := leftBoarder(arr2, max)
		if rIdx-lIdx <= 1 {
			cnt++
		}
	}

	return cnt
}
