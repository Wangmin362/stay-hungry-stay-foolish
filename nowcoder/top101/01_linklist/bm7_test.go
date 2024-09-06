package _1_linklist

import "testing"

// https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != pHead {
				slow = slow.Next
				pHead = pHead.Next
			}
			return slow
		}
	}
	return nil
}

func TestXxx(t *testing.T) {
	tmp := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	tmp.Next.Next.Next = tmp
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: tmp}}
	EntryNodeOfLoop(l1)
}
