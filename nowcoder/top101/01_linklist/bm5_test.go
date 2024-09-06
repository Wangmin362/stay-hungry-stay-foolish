package _1_linklist

import "container/heap"

// https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6

// 最小堆解题
func mergeKLists01(lists []*ListNode) *ListNode {
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
		if n.Next != nil {
			heap.Push(&h, n.Next)
		}
		curr.Next = n
		curr = curr.Next
	}

	return dummy.Next
}

type hp []*ListNode

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(x any)         { *h = append(*h, x.(*ListNode)) }
func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
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

// 分支
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

// 每次获取最小节点
func mergeKLists03(lists []*ListNode) *ListNode {
	getMinNode := func() *ListNode {
		var minNode *ListNode
		for _, head := range lists {
			if head == nil {
				continue
			}
			if minNode == nil {
				minNode = head
			} else if head.Val < minNode.Val {
				minNode = head
			}
		}

		if minNode == nil {
			return minNode
		}

		// 移动最小节点所在的链表，方便下一次做对比
		for idx, head := range lists {
			if head == minNode {
				lists[idx] = lists[idx].Next
				break
			}
		}

		return minNode
	}

	dummy := &ListNode{}
	curr := dummy
	for no := getMinNode(); no != nil; {
		curr.Next = no
		curr = curr.Next
		no = getMinNode()
	}

	return dummy.Next
}
