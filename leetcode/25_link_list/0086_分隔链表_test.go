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

func partition02(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var lef *ListNode // 第一个小于x的节点
	var gef *ListNode // 第一个大于x的节点
	var le *ListNode  // 小于x节点链
	var ge *ListNode  // 大于x节点链

	dummy := &ListNode{Next: head}

	for head != nil {
		if head.Val >= x {
			if gef == nil {
				gef = head
				ge = head
			} else {
				ge.Next = head
				ge = ge.Next
			}
		} else {
			if lef == nil {
				lef = head
				le = head
			} else {
				le.Next = head
				le = le.Next
			}
		}
		head = head.Next
	}
	if gef != nil {
		ge.Next = nil // 断开链表
	}

	if lef != nil {
		le.Next = gef
		dummy.Next = lef
	}

	return dummy.Next
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
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1}}, target: 2, expect: &ListNode{Val: 1, Next: &ListNode{Val: 1}}},
		{head: nil, target: 8, expect: nil},
	}

	for _, test := range testdata {
		get := partition02(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.target, expect, get)
		}
	}

}
