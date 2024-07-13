package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/linked-list-cycle-ii/
// 视频讲解：https://www.bilibili.com/video/BV1if4y1d7ob/

// 解题思路：使用快慢指针，满指针一次走一个节点，快指针一次走两个节点。如果快满指针相遇，那么一定有环。此时让慢指针继续走，然后从链表头再
// 用一个指针以一次一个节点的速度走，直到相遇，说明找到了环的入口点。

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	c1 := &ListNode{Val: 3}
	c2 := &ListNode{Val: 2}
	c3 := &ListNode{Val: 0}
	c4 := &ListNode{Val: -4}
	c1.Next = c2
	c2.Next = c3
	c3.Next = c4
	c4.Next = c2
	testdata := []struct {
		head   *ListNode
		expect *ListNode
	}{
		//{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		//	expect: nil,
		//},
		{head: c1, expect: c2},
		{head: nil, expect: nil},
	}

	for _, test := range testdata {
		get := detectCycle(test.head)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("expect:%v, get:%v", expect, get)
		}
	}
}
