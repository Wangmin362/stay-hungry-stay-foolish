package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/palindrome-linked-list/description/

// 解法一，非常简单，直接把链表中的元素放入到数组中，然后判断数组是否是回文即可
func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 解法二，第一次遍历计算链表的长度，然后翻转前半段节点，最后使用快慢指针对比，然后还原链表
// TODO
func isPalindrome02(head *ListNode) bool {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length < 2 { // 快速判断
		return true
	}

	mid := length / 2
	fast := mid
	if length%2 == 1 {
		fast += 2
	}

	//reverseListFunc : = func(n int) { // 反转链表的前N个节点
	//	var prev *ListNode
	//	curr := head
	//	for n > 0 {
	//		curr
	//	}
	//}

	return true
}

func TestIsPalindrome(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		expect bool
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: false,
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
			expect: true,
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
			expect: true,
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 3}},
			expect: true,
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2}},
			expect: false,
		},
		{head: &ListNode{Val: 3},
			expect: true,
		},
		{head: nil,
			expect: true,
		},
	}
	for _, test := range testdata {
		get := isPalindrome(test.head)
		if get != test.expect {
			t.Fatalf("list:%v, expect:%v, get:%v", test.head, test.expect, get)
		}
	}
}
