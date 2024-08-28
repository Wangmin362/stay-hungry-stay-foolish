package _9_binary_search

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/description/

// 题目分析：由于是向上取整，因此所有元素处于选择的数字d之后，最小的情况下就是len(nums)，因为是向上取整
// 所以，我们选择的数字的大小范围在[1, max(nums)]之间，根据除法的特性，当选择的数字d增加时，结果是减小的，
// 因此total也是减小的，若d减小时，total则是单调递增的

// 题目要求计算小于等于threshold结果的最小值，选取的d在[1, max(nums)]之间，因此total的值时从大到小单调
// 递减的，越往右，total越小，但是选取的数字d越大，因此一定存在一个最小的数字d刚好total小于threshold，并且
// d+1的total是大于threshold的

func smallestDivisor(nums []int, threshold int) int {
	sum := func(d int) int {
		var total int
		for _, n := range nums {
			total += int(math.Ceil(float64(n) / float64(d)))
		}
		return total
	}

	leftBoarder := func() int {
		left, right := 1, slices.Max(nums)
		for left <= right {
			mid := left + (right-left)>>1
			total := sum(mid)
			if total <= threshold {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return right + 1
	}
	return leftBoarder()
}
