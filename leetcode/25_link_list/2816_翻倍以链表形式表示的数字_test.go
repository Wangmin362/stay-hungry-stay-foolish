package _1_array

import (
	"fmt"
	"testing"
)

func doubleIt(head *ListNode) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			nxt := head.Next
			head.Next = prev
			prev = head
			head = nxt
		}
		return prev
	}

	head = reverse(head)
	overflow := 0
	curr := head
	var prev *ListNode
	for curr != nil {
		product := 2*curr.Val + overflow
		if product > 9 {
			overflow = product / 10
			product %= 10
		} else {
			overflow = 0
		}
		curr.Val = product
		prev = curr
		curr = curr.Next
	}
	if overflow > 0 {
		prev.Next = &ListNode{Val: overflow}
	}

	return reverse(head)
}

func TestDoubleIt(t *testing.T) {
	l := &ListNode{1, &ListNode{5, &ListNode{9, nil}}}
	l = doubleIt(l)
	fmt.Println(l)

	l = &ListNode{9, &ListNode{8, &ListNode{9, nil}}}
	l = doubleIt(l)
	fmt.Println(l)
}
