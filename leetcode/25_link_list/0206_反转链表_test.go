package _1_array

import (
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil

	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func TestReverseList(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 3},
			expect: &ListNode{Val: 3},
		},
	}
	for _, test := range testdata {
		get := reverseList(test.head)
		expect := test.expect
		if linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}
}
