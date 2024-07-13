package _1_array

import (
	"testing"
)

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	getLen := func(head *ListNode) int {
		len := 0
		curr := head
		for curr != nil {
			curr = curr.Next
			len++
		}
		return len
	}

	aLen := getLen(headA)
	bLen := getLen(headB)

	if aLen > bLen {
		for i := 0; i < aLen-bLen; i++ {
			headA = headA.Next
		}
		for headA != nil && headB != nil {
			if headA == headB {
				return headA
			}
			headA = headA.Next
			headB = headB.Next
		}
		return nil
	} else {
		for i := 0; i < bLen-aLen; i++ {
			headB = headB.Next
		}
		for headA != nil && headB != nil {
			if headA == headB {
				return headA
			}
			headA = headA.Next
			headB = headB.Next
		}
		return nil
	}

	return nil
}

func TestGetIntersectionNode(t *testing.T) {
	testdata := []struct {
		head1  *ListNode
		head2  *ListNode
		expect *ListNode
	}{
		//{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		//	expect: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
		//},
		//{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		//	expect: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
		//},
		//{head: &ListNode{Val: 3, Next: &ListNode{Val: 8}}, expect: &ListNode{Val: 8, Next: &ListNode{Val: 3}}},
		//{head: &ListNode{Val: 3}, expect: &ListNode{Val: 3}},
		//{head: nil, expect: nil},
	}

	for _, test := range testdata {
		get := getIntersectionNode(test.head1, test.head2)
		expect := test.expect
		if linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}
}
