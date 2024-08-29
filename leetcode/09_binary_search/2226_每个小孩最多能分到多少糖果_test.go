package _9_binary_search

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/maximum-candies-allocated-to-k-children/description/

// 题目分析：每个小孩分配t个糖果，那么总共可以分配total = sum(candies[i]/t),只有当total大于等于k
// 的时候，才能正确分配，显然我们需要找到t是合法，并且t+1无法分配，此时t就是那个值
// 搜索区间为[1, max(candies)]

func maximumCandies(candies []int, k int64) int {
	var sumCandies int
	for _, c := range candies {
		sumCandies += c
	}
	if sumCandies < int(k) { // 小孩的人数比糖果总数还多，肯定不够分
		return 0
	}

	sumK := func(t int) int { // 每个小孩分配的糖果为t
		var sum int
		for _, c := range candies {
			sum += c / t
		}
		return sum
	}

	left, right := 1, slices.Max(candies)
	for left <= right {
		mid := left + (right-left)>>1
		sumKid := sumK(mid)
		if sumKid >= int(k) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func TestName(t *testing.T) {
	fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
}
