package _1_array

import (
	"fmt"
	"testing"
)

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	insert := func(head, node *ListNode) *ListNode {
		dummy := &ListNode{Next: head}
		curr := dummy
		for curr.Next != nil {
			if curr.Next.Val >= node.Val {
				node.Next = curr.Next
				curr.Next = node
				return dummy.Next
			}
			curr = curr.Next
		}
		curr.Next = node

		return dummy.Next
	}

	ordered := head
	head = head.Next
	ordered.Next = nil

	for head != nil {
		nxt := head.Next
		node := head
		node.Next = nil
		head = nxt
		ordered = insert(ordered, node)
	}

	return ordered
}

func TestInsertionSortList(t *testing.T) {
	l := &ListNode{4, &ListNode{2, &ListNode{9, &ListNode{0, &ListNode{3, nil}}}}}
	l = insertionSortList(l)
	fmt.Println(l)

	l = &ListNode{4, &ListNode{2, nil}}
	l = insertionSortList(l)
	fmt.Println(l)

	l = &ListNode{0, &ListNode{2, nil}}
	l = insertionSortList(l)
	fmt.Println(l)
}
