package _1_array

import (
	"testing"
)

// 题目：https://leetcode.cn/problems/rotate-list/description/

// 解题思路：使用快慢指针，快指针先走k步，然后满指针一起走，快指针为nil时，旋转即可

func rotateRight(head *ListNode, k int) *ListNode {
	curr := head
	length := 0
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length <= 0 {
		return head
	}

	modK := k % length // 可能移动的次数比链表节点多，因此只需要移动余数次就可以了
	if modK == 0 {
		return head
	}
	slow := head
	fast := head
	preFast := head // 记录fast的前面一个节点
	for i := modK; i >= 0; i-- {
		fast = fast.Next
		if i > 0 {
			preFast = preFast.Next
		}
	}
	for fast != nil {
		slow = slow.Next
		preFast = preFast.Next
		fast = fast.Next
	}

	tmpHead := head
	head = slow.Next
	slow.Next = nil
	preFast.Next = tmpHead
	return head
}

func TestRotateRight(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		target int
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 0,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 1,
			expect: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 2,
			expect: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 3,
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 4,
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 5,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 6,
			expect: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}},
		},
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			target: 7,
			expect: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{head: nil, target: 8, expect: nil},
	}

	for _, test := range testdata {
		get := rotateRight(test.head, test.target)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.target, expect, get)
		}
	}

}
