package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/?envType=daily-question&envId=2024-09-09

func mergeNodes(head *ListNode) *ListNode {
	prev := head.Next
	fast := prev.Next
	for fast != nil {
		if fast.Val != 0 {
			prev.Val += fast.Val
			fast = fast.Next
		} else {
			prev.Next = fast.Next
			prev = prev.Next
			if prev != nil {
				fast = prev.Next
			} else {
				fast = nil
			}
		}
	}

	return head.Next
}

func TestMergeNodes(t *testing.T) {
	l1 := &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}}
	fmt.Println(mergeNodes(l1))

}
