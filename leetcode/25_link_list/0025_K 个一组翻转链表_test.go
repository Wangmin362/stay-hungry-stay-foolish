package _1_array

import "testing"

// 题目：https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

func reverseKGroup01(head *ListNode, k int) *ListNode {
	getLen := func(node *ListNode) int {
		var res int
		for node != nil {
			node = node.Next
			res++
		}
		return res
	}

	length := getLen(head)
	if length <= 1 || k <= 1 || length < k {
		return head
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	for length >= k {
		step := k
		curr := p0.Next
		nxtP0 := p0.Next
		var prev *ListNode
		for step > 0 {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
			step--
		}
		p0.Next.Next = curr
		p0.Next = prev
		p0 = nxtP0
		length -= k
	}

	return dummy.Next
}

func TestReverseKGroup01(t *testing.T) {
	var testdata = []struct {
		head   *ListNode
		k      int
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:      2,
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}},
		},
		{head: nil, k: 8, expect: nil},
	}

	for _, test := range testdata {
		get := reverseKGroup01(test.head, test.k)
		expect := test.expect
		if !linkListEqual(get, expect) {
			t.Fatalf("tar:%d,expect:%v, get:%v", test.k, expect, get)
		}
	}

}
