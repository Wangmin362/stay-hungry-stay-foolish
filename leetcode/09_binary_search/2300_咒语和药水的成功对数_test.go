package _9_binary_search

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/description/

// 题目分析：其实就是找出spells[i]*points[j] >= success的数目

// 思路一：直接两层for循环遍历，O(n * m)
// 思路二：potions排序，然后O(n *log(m))

// 思路一：两层for循环
func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, 0, len(spells))
	for _, sp := range spells {
		var cnt int
		for _, po := range potions {
			if int64(sp*po) >= success {
				cnt++
			}
		}
		res = append(res, cnt)
	}
	return res
}

// 思路二：
func successfulPairs02(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
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
	res := make([]int, 0, len(spells))
	for _, sp := range spells {
		target := int(math.Ceil(float64(success) / float64(sp)))
		idx := leftBoarder(potions, target)
		cnt := len(potions) - idx
		res = append(res, cnt)
	}
	return res
}
