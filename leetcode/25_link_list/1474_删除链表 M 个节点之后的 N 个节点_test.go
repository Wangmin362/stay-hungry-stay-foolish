package _1_array

import (
	"fmt"
	"testing"
)

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	slow, fast := head, head
	remain, del := 1, 0 // 第一个节点肯定是要保留的
	for fast != nil {
		if remain < m {
			slow = slow.Next
			fast = fast.Next
			remain++
			continue
		}
		if del < n {
			fast = fast.Next
			del++
			continue
		}
		slow.Next = fast.Next
		slow = slow.Next
		fast = fast.Next
		remain = 1
		del = 0
	}
	if slow != nil {
		slow.Next = fast
	}

	return head
}

func TestName(t *testing.T) {
	//l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6,
	//	&ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, &ListNode{11,
	//		&ListNode{12, &ListNode{13, nil}}}}}}}}}}}}}
	//l = deleteNodes(l, 2, 3)
	//fmt.Println(l)

	l2 := &ListNode{6, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{2,
		&ListNode{8, &ListNode{9, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}}}}}}}
	l2 = deleteNodes(l2, 2, 1)
	fmt.Println(l2)
}
