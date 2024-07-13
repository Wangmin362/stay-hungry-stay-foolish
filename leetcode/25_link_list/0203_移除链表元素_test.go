package _1_array

import (
	"testing"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

func removeElements01(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

func TestRemoveElement(t *testing.T) {
	testdata := []struct {
		head   *ListNode
		target int
		expect *ListNode
	}{
		{
			head:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 3,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		},
		{
			head:   &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 3,
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 4}},
		},
		{
			head:   &ListNode{Val: 3},
			target: 3,
			expect: nil,
		},
	}

	for _, test := range testdata {
		get := removeElements01(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}
}
