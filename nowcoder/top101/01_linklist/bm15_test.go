package _1_linklist

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			slow.Next = fast.Next
		} else {
			slow = fast
		}
		fast = fast.Next
	}

	return head
}
func TestDeleteDuplicates(t *testing.T) {
	l := &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 5}}}}}
	fmt.Println(deleteDuplicates(l))
}
