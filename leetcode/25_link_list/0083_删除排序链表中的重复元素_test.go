package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

// 解题思路：使用快慢指针，只要快指针和前一个节点的值相同，就需要移除快指针值向的节点，否则通过是移动，快慢节点

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next
	for fast != nil {
		if fast.Val == slow.Val {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
			expect: &ListNode{Val: 1},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 3}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
		{head: nil, expect: nil},
	}

	for _, test := range testdata {
		get := deleteDuplicates(test.head)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
