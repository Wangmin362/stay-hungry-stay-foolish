package _0_basic

import "container/heap"

// https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

type hpp []int

func (h *hpp) Len() int { return len(*h) }
func (h *hpp) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *hpp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hpp) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *hpp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxProduct(nums []int) int {
	h := hpp(nums)
	heap.Init(&h)

	x1 := heap.Pop(&h).(int) - 1
	x2 := heap.Pop(&h).(int) - 1
	return x1 * x2
}
