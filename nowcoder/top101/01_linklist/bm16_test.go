package _1_linklist

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev, slow, fast := dummy, head, head.Next
	for fast != nil {
		if slow.Val != fast.Val {
			prev = prev.Next
			slow = fast
			fast = fast.Next
		} else {
			for fast != nil && fast.Val == slow.Val {
				fast = fast.Next
			}

			if fast == nil {
				prev.Next = nil
			} else {
				prev.Next = fast
				slow = fast
				fast = fast.Next
			}
		}
	}

	return dummy.Next
}
func TestDeleteDuplicatesII(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1,
		Next: &ListNode{Val: 1}}}}}
	fmt.Println(deleteDuplicatesII(l))
}
