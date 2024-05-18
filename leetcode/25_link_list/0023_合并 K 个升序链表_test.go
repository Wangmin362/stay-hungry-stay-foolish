package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/merge-k-sorted-lists/description/

// 解题思路：写一个for循环，对比所有链表的节点值，谁小就选谁，同时移动一格

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	getMinNode := func() *ListNode {
		var tmpVal *ListNode
		for _, head := range lists {
			if tmpVal == nil {
				tmpVal = head
			} else {
				if head != nil && head.Val < tmpVal.Val {
					tmpVal = head // 此时还不能移动节点，因为它不一定是最小的
				}
			}
		}
		if tmpVal != nil {
			for idx := 0; idx < len(lists); idx++ {
				if lists[idx] == tmpVal {
					lists[idx] = lists[idx].Next // 移动节点
					break
				}
			}
		}

		return tmpVal
	}

	curr := dummy
	for minNode := getMinNode(); minNode != nil; {
		curr.Next = minNode
		curr = curr.Next
		minNode = getMinNode()
	}

	return dummy.Next
}

func TestMergeKLists(t *testing.T) {
	var testdata = []struct {
		lists  []*ListNode
		expect *ListNode
	}{
		{
			lists: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				&ListNode{Val: 2, Next: &ListNode{Val: 6}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}},
		},
		{
			lists: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				nil,
				&ListNode{Val: 2, Next: &ListNode{Val: 6}}},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}},
		},
	}

	for _, test := range testdata {
		get := mergeKLists(test.lists)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}

}
