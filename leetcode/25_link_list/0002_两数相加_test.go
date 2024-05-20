package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	jw := 0
	curr := dummy
	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2 + jw
		if sum > 9 {
			jw = 1
		} else {
			jw = 0
		}
		v := sum % 10
		curr.Next = &ListNode{Val: v}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if jw > 0 {
		curr.Next = &ListNode{Val: jw}
	}

	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	var testdata = []struct {
		head1  *ListNode
		head2  *ListNode
		expect *ListNode
	}{
		{head1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			head2:  &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expect: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{head1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			head2:  &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			expect: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
		{head1: &ListNode{Val: 0},
			head2:  &ListNode{Val: 0},
			expect: &ListNode{Val: 0},
		},
		{head1: nil, head2: nil, expect: nil},
	}

	for _, test := range testdata {
		get := addTwoNumbers(test.head1, test.head2)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("head1:%v, head2:%v, expect:%v, get:%v", test.head1, test.head2, expect, get)
		}
	}

}
