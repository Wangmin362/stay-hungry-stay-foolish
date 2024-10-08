package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/reverse-linked-list-ii/

// 解法一，把链表放入到数组，然后交换，最后重新构造链表

func reverseBetween01(head *ListNode, left int, right int) *ListNode {
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	left--
	right--
	for left < right && left < len(arr) && right < len(arr) {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	dummy := &ListNode{}
	curr := dummy
	for _, node := range arr {
		curr.Next = node
		curr = curr.Next
	}
	curr.Next = nil

	return dummy.Next
}

// 解法二，本质上就是一个反转链表，只不过只反转中间的一部分

func reverseBetween02(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	step := left
	for step > 1 {
		step--
		p0 = p0.Next
	}
	step = right - left + 1
	var pre *ListNode
	cur := p0.Next
	for step > 0 {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		step--
	}
	p0.Next.Next = cur
	p0.Next = pre

	return dummy.Next
}

func TestReverseBetween(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		left   int
		right  int
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			left:   2,
			right:  4,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}}},
		},
		{head: nil, left: 8, right: 10, expect: nil},
	}

	for _, test := range testdata {
		get := reverseBetween02(test.head, test.left, test.right)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("left:%v, right:%v, expect:%v, get:%v", test.left, test.right, test.expect, get)
		}
	}

}
