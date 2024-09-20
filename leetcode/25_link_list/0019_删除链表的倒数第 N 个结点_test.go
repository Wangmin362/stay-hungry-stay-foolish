package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// 解题思路：只需要使用dummy节点进行统一，并且使用快慢指针，然后画图即可解决此问题

// 使用快慢指针
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

func removeNthFromEnd03(head *ListNode, n int) *ListNode {
	getLen := func(head *ListNode) int {
		var res int
		for head != nil {
			res++
			head = head.Next
		}
		return res
	}

	length := getLen(head)
	if n <= 0 || n > length {
		return head
	}

	dummy := &ListNode{Next: head}
	step := length - n
	curr := dummy
	for step > 0 {
		curr = curr.Next
		step--
	}

	curr.Next = curr.Next.Next

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
		get := removeNthFromEnd03(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.target, expect, get)
		}
	}
}
