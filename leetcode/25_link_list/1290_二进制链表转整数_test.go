package _1_array

import "testing"

// 题目：https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/description/

func getDecimalValue(head *ListNode) int {
	getLen := func(node *ListNode) int {
		var res int
		for node != nil {
			res++
			node = node.Next
		}
		return res
	}

	length := getLen(head) - 1
	var res int
	for head != nil {
		res += head.Val * (1 << length)
		head = head.Next
		length--
	}
	return res
}

func TestGetDecimalValue(t *testing.T) {
	var testdata = []struct {
		head1  *ListNode
		expect int
	}{
		{head1: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
			expect: 5,
		},
	}

	for _, test := range testdata {
		get := getDecimalValue(test.head1)
		expect := test.expect
		if get != expect {
			t.Fatalf("head1:%v, expect:%v, get:%v", test.head1, expect, get)
		}
	}

}
