package _0_basic

import (
	"container/heap"
	"math"
)

// https://leetcode.cn/problems/take-gifts-from-the-richest-pile/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

// 最大堆
func pickGifts(gifts []int, k int) int64 {
	h := hpp(gifts)
	heap.Init(&h)

	for i := 0; i < k; i++ {
		x := heap.Pop(&h).(int)
		heap.Push(&h, int(math.Sqrt(float64(x))))
	}

	var res int
	for h.Len() > 0 {
		x := heap.Pop(&h).(int)
		res += x
	}
	return int64(res)
}
