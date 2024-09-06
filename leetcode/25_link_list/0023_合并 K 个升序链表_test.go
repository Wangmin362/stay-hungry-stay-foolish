package _1_array

import (
	"container/heap"
	"testing"
)

// 题目：https://leetcode.cn/problems/merge-k-sorted-lists/description/

// 解题思路：写一个for循环，对比所有链表的节点值，谁小就选谁，同时移动一格

func mergeKLists00(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	getMinNode := func() *ListNode {
		var tmpVal *ListNode
		for _, head := range lists {
			if tmpVal == nil {
				tmpVal = head
			} else {
				if head != nil && head.Val < tmpVal.Val {
					tmpVal = head // 此时还不能移动节点，因为它不一定是最小的
				}
			}
		}
		if tmpVal != nil {
			for idx := 0; idx < len(lists); idx++ {
				if lists[idx] == tmpVal {
					lists[idx] = lists[idx].Next // 移动节点
					break
				}
			}
		}

		return tmpVal
	}

	curr := dummy
	for minNode := getMinNode(); minNode != nil; {
		curr.Next = minNode
		curr = curr.Next
		minNode = getMinNode()
	}

	return dummy.Next
}

// 解法二：最小堆，每次把所有可能的节点放入最小堆当中，并且每次获取堆顶元素构建链表
type hp []*ListNode

func (h *hp) Len() int { return len(*h) }

func (h *hp) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val } // 最小堆

func (h *hp) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *hp) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func mergeKLists0906MinHeap(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}

	// 堆化
	heap.Init(&h)
	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		n := heap.Pop(&h).(*ListNode)
		if n.Next != nil { // 下一个可能的最小值还有可能在同一个链表当中
			heap.Push(&h, n.Next)
		}
		curr.Next = n
		curr = curr.Next
	}

	return dummy.Next
}

func mergeTwoList(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for h1 != nil && h2 != nil {
		if h1.Val > h2.Val {
			curr.Next = h2
			h2 = h2.Next
		} else {
			curr.Next = h1
			h1 = h1.Next
		}

		curr = curr.Next
	}
	if h1 != nil {
		curr.Next = h1
	}
	if h2 != nil {
		curr.Next = h2
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return mergeTwoList(left, right)
}

func TestMergeKLists(t *testing.T) {
	var testdata = []struct {
		lists  []*ListNode
		expect *ListNode
	}{
		{
			lists: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				&ListNode{Val: 2, Next: &ListNode{Val: 6}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}},
		},
		{
			lists: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				nil,
				&ListNode{Val: 2, Next: &ListNode{Val: 6}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}},
		},
	}

	for _, test := range testdata {
		get := mergeKLists(test.lists)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
