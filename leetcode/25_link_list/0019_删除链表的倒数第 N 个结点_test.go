package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// 解题思路：只需要使用dummy节点进行统一，并且使用快慢指针，然后画图即可解决此问题

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	dummy := &ListNode{Next: head}

	curr := dummy
	slow := curr
	fast := curr
	for n >= 0 && fast != nil {
		fast = fast.Next
		n--
	}
	if n >= 0 { // 如果n还是大于等于0，说明链表根本就没有n那么长，直接返回原始链表即可
		return dummy.Next
	}

	// 到了这里，fast指针已经比slow指针快了n步骤，此时只需要同时移动slow, fast，一旦fast=nil，此时只需要slow.next = slow.next.next即可
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

func removeNthFromEnd02(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	dummy := &ListNode{Next: head}

	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ { // 快指针先走N步
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	testdata := []struct {
		head   *ListNode
		target int
		expect *ListNode
	}{
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 1,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 2,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 3,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 4,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 5,
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 6,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		},
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 0,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		},
		{head: nil, target: 8, expect: nil},
	}

	for _, test := range testdata {
		get := removeNthFromEnd02(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.target, expect, get)
		}
	}
}
