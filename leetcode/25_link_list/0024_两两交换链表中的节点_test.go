package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/swap-nodes-in-pairs/description/

// 解题思路：只需要使用dummy节点进行统一，然后画图即可解决此问题

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil { // 当前指针的必须要有两个节点，否则不需要进行节点交换，直接退出即可
		slow := curr.Next
		fast := slow.Next
		tmp := fast.Next

		slow.Next = tmp
		fast.Next = slow
		curr.Next = fast

		curr = slow // 移动当前指针
	}

	return dummy.Next
}

func TestSwapPairs(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 8}}, expect: &ListNode{Val: 8, Next: &ListNode{Val: 3}}},
		{head: &ListNode{Val: 3}, expect: &ListNode{Val: 3}},
		{head: nil, expect: nil},
	}

	for _, test := range testdata {
		get := swapPairs(test.head)
		expect := test.expect
		if linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
