package _0_basic

import (
	"container/heap"
	"fmt"
	"testing"
)

type hhp []int

func (h *hhp) Len() int { return len(*h) }

func (h *hhp) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *hhp) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *hhp) Push(x any) { *h = append(*h, x.(int)) }

func (h *hhp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := hhp(nums)
	heap.Init(&h)

	for h.Len() > k {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
