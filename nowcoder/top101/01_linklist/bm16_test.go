package _1_linklist

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	preSlow := dummy
	slow := head
	fast := head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			preSlow = preSlow.Next
			slow = slow.Next
			fast = fast.Next
		} else {
			for fast != nil && fast.Val == slow.Val {
				fast = fast.Next
			}
			if fast == nil {
				preSlow.Next = nil
			} else {
				preSlow.Next = fast
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
