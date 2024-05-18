package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/partition-list/

func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	biggerEqual := &ListNode{}

	curr := head
	lessCurr := less
	biggerEqualCurr := biggerEqual
	for curr != nil {
		if curr.Val < x {
			lessCurr.Next = curr
			lessCurr = lessCurr.Next
		} else {
			biggerEqualCurr.Next = curr
			biggerEqualCurr = biggerEqualCurr.Next
		}
		curr = curr.Next
	}
	lessCurr.Next = biggerEqual.Next
	biggerEqualCurr.Next = nil

	return less.Next
}

func TestPartition(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		target int
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}},
			target: 3,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
		},
		{head: nil, target: 8, expect: nil},
	}

	for _, test := range testdata {
		get := partition(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.target, expect, get)
		}
	}

}
