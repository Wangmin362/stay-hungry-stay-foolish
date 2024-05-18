package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/merge-two-sorted-lists/description/

// 解题思路：对比两个节点头，谁小选择谁

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}

	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr.Next = list2 // 谁小选择谁
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}

func TestMergeTwoLists(t *testing.T) {
	var testdata = []struct {
		head1  *ListNode
		head2  *ListNode
		expect *ListNode
	}{
		{head1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			head2:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}}},
		},
		{head1: nil, head2: nil, expect: nil},
	}

	for _, test := range testdata {
		get := mergeTwoLists(test.head1, test.head2)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("head1:%v, head2:%v, expect:%v, get:%v", test.head1, test.head2, expect, get)
		}
	}

}
