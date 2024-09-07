package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(x any)         { *h = append(*h, x.(int)) }
func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	h := hp{4, 6, 9, 3, 6}
	heap.Init(&h)
	for len(h) > 0 {
		x := heap.Pop(&h)
		fmt.Println(x)
	}
}
