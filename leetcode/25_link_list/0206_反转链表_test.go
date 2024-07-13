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

// 快慢指针
func reverseList01(head *ListNode) *ListNode {
	var dummy *ListNode
	slow, fast := dummy, head
	for fast != nil {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}

	return slow
}

func TestReverseList(t *testing.T) {
	testdata := []struct {
		head   *ListNode
		expect *ListNode
	}{
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
		{
			head:   &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			head:   &ListNode{Val: 3},
			expect: &ListNode{Val: 3},
		},
	}
	for _, test := range testdata {
		get := reverseList01(test.head)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}
}
