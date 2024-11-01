package _1_array

import (
	"fmt"
	"testing"
)

//https://leetcode.cn/problems/reorder-list/?envType=problem-list-v2&envId=linked-list&status=TO_DO

// 解题思路：非常简单，其实就是找到链表的中点，然后翻转后面的链表，然后先从前面的链表取值，再从后面的链表取值
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil { // 如果链表少于两个节点，直接返回
		return
	}

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

	slow, fast := head, head.Next.Next // 找到链表的中点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1 := head
	l2 := reverse(slow.Next)
	slow.Next = nil // 断开链

	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		curr.Next = l1
		l1 = l1.Next
		curr = curr.Next
		curr.Next = l2
		l2 = l2.Next
		curr = curr.Next
	}

	if l2 != nil {
		curr.Next = l2
	}

	head = dummy.Next
}

func TestReorderList(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(l1)
	fmt.Println(l1)

	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(l2)
	fmt.Println(l2)
}
