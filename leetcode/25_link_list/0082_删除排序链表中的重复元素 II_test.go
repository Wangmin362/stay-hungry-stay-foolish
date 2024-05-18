package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

func deleteDuplicatesTwo(head *ListNode) *ListNode {
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

func TestDeleteDuplicatesTwo(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}},
			expect: &ListNode{Val: 1},
		},
		{head: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}},
			expect: nil,
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 8}}, expect: &ListNode{Val: 3, Next: &ListNode{Val: 8}}},
		{head: &ListNode{Val: 3}, expect: &ListNode{Val: 3}},
		{head: nil, expect: nil},
	}

	for _, test := range testdata {
		get := deleteDuplicatesTwo(test.head)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
