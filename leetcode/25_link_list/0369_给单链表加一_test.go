package _1_array

import (
	"fmt"
	"testing"
)

func plusOne(head *ListNode) *ListNode {
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
	curr := head
	var prev *ListNode
	overflow := 1
	for curr != nil {
		sum := curr.Val + overflow
		if sum > 9 {
			overflow = 1
			sum %= 10
		} else {
			overflow = 0
		}
		curr.Val = sum
		prev = curr
		curr = curr.Next
	}
	if overflow == 1 {
		prev.Next = &ListNode{Val: 1}
	}

	return reverse(head)
}

func TestPlusOne(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l = plusOne(l)
	fmt.Println(l)

	l = &ListNode{1, &ListNode{2, &ListNode{9, nil}}}
	l = plusOne(l)
	fmt.Println(l)

	l = &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	l = plusOne(l)
	fmt.Println(l)
}
