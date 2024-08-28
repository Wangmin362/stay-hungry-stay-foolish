package _9_binary_search

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/koko-eating-bananas/description/

// 题目分析，其实就是一个向上取整的问题，根据题目表述，由于每次只能吃一堆，因此至少需要len(piles)个小时才能
// 把所有的香蕉吃完，因此h必须大于等于len(piles)，否则题目没有意义，或者说是无解。 吃香蕉的个数最小为1，最大
// 为max(piles)，即保证每个小时可以吃一堆香蕉

func minEatingSpeed(piles []int, h int) int {
	if h < len(piles) { // 这种情况下，肯定吃不完，无解
		return -1
	}

	sum := func(k int) int { // 每次吃k个相加，返回一共需要多少个小时可以吃完
		var res int
		for _, pi := range piles {
			res += int(math.Ceil(float64(pi) / float64(k)))
		}
		return res
	}

	left, right := 1, slices.Max(piles)
	for left <= right {
		mid := left + (right-left)>>1
		totalHours := sum(mid)
		if totalHours <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
