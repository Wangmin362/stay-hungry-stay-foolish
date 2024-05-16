package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/intersection-of-two-linked-lists/

// 解题思路：只需要使用dummy节点进行统一，然后画图即可解决此问题

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return nil
}

func TestGetIntersectionNode(t *testing.T) {
	var testdata = []struct {
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
