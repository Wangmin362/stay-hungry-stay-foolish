package _1_array

import "testing"

// 题目：https://leetcode.cn/problems/add-two-numbers-ii/description/

// 解法一：先反转两个链表，然后从最低为开始相加
func addTwoNumbers01(l1 *ListNode, l2 *ListNode) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var prev *ListNode
		for node != nil {
			tmp := node.Next
			node.Next = prev
			prev = node
			node = tmp
		}
		return prev
	}

	l1 = reverse(l1)
	l2 = reverse(l2)
	dummy := &ListNode{}
	curr := dummy
	mod := 0
	for l1 != nil || l2 != nil {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sum := l1Val + l2Val + mod
		if sum > 9 {
			sum = sum % 10
			mod = 1
		} else {
			mod = 0
		}

		if l2 == nil {
			curr.Next = l1
			curr.Next.Val = sum
			curr = curr.Next
			l1 = l1.Next
		} else if l1 == nil {
			curr.Next = l2
			curr.Next.Val = sum
			curr = curr.Next
			l2 = l2.Next
		} else {
			curr.Next = l1
			curr.Next.Val = sum
			curr = curr.Next
			l1 = l1.Next
			l2 = l2.Next
		}

	}
	if mod == 1 {
		curr.Next = &ListNode{Val: 1}
	}

	return reverse(dummy.Next)
}

func TestAddTwoNumbers01(t *testing.T) {
	var testdata = []struct {
		head1  *ListNode
		head2  *ListNode
		expect *ListNode
	}{
		{head1: &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
			head2:  &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expect: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 7}}}},
		},
		{head1: &ListNode{Val: 5},
			head2:  &ListNode{Val: 5},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 0}},
		},
		{head1: nil, head2: nil, expect: nil},
	}

	for _, test := range testdata {
		get := addTwoNumbers01(test.head1, test.head2)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("head1:%v, head2:%v, expect:%v, get:%v", test.head1, test.head2, expect, get)
		}
	}

}
