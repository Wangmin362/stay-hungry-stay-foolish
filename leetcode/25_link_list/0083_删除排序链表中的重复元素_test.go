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

// 这种方式会导致内存泄露，因为被删除的节点没有释放内存
func deleteDuplicates02(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
		}
		fast = fast.Next
	}
	slow.Next = fast
	return head
}

// 释放被删除节点的内存空间，
func deleteDuplicatesGC(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		} else {
			nxt := fast.Next
			fast.Next = nil // 释放内存空间
			fast = nxt
		}
	}
	slow.Next = fast
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
		get := deleteDuplicatesGC(test.head)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
