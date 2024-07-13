package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/linked-list-cycle/description/

// 解题思路：使用快慢指针，慢指针步进为一，快指针步进为二，只要快慢指针相等，即为有环

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{Next: head}

	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func hasCycle01(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func TestHasCycle(t *testing.T) {
	testdata := []struct {
		head   *ListNode
		expect bool
	}{
		{head: nil, expect: false},
	}

	for _, test := range testdata {
		get := hasCycle(test.head)
		if get != test.expect {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
