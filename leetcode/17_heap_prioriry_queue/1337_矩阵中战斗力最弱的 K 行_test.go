package _0_basic

import (
	"container/heap"
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

func kWeakestRows(mat [][]int, k int) []int {
	type peer struct {
		row int
		cap int
	}
	ps := make([]peer, len(mat))
	for idx, ma := range mat {
		ps[idx] = peer{row: idx}
		for _, m := range ma {
			if m == 0 {
				break
			}
			ps[idx].cap++
		}
	}
	slices.SortFunc(ps, func(a, b peer) int {
		if a.cap > b.cap {
			return 1
		}
		if a.cap == b.cap {
			if a.row > b.row {
				return 1
			} else {
				return -1
			}
		}
		return -1
	})

	res := make([]int, k)
	for idx, p := range ps {
		if idx >= k {
			break
		}
		res[idx] = p.row
	}
	return res
}

type peer struct {
	row int
	cap int
}
type hp []peer

func (h *hp) Len() int { return len(*h) }

func (h *hp) Less(i, j int) bool {
	if (*h)[i].cap == (*h)[j].cap {
		return (*h)[i].row < (*h)[j].row
	}
	return (*h)[i].cap < (*h)[j].cap
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(x any) { *h = append(*h, x.(peer)) }

func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func kWeakestRows02(mat [][]int, k int) []int {
	getCap := func(arr []int) int {
		left, right := 0, len(arr)-1 // [left, right]
		for left <= right {
			mid := left + (right-left)>>1
			if arr[mid] == 0 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return left
	}

	h := make(hp, len(mat))
	for idx, ma := range mat {
		h[idx] = peer{
			row: idx,
			cap: getCap(ma),
		}
	}
	heap.Init(&h)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		x := heap.Pop(&h).(peer)
		res[i] = x.row
	}

	return res
}

func TestKWeakestRows(t *testing.T) {
	mat := [][]int{
		{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0},
	}
	fmt.Println(kWeakestRows02(mat, 2))
}
