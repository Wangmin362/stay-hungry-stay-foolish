package _1_array

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (r *ListNode) String() string {
	res := ""
	tmp := r
	for tmp != nil {
		res += fmt.Sprintf("%d -> ", tmp.Val)
		tmp = tmp.Next
	}
	res += "nil"
	return res
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
func TestRemoveElement(t *testing.T) {
	var twoSumTest = []struct {
		head   *ListNode
		target int
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 3,
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			target: 3,
			expect: &ListNode{Val: 2, Next: &ListNode{Val: 4}},
		},
		{head: &ListNode{Val: 3},
			target: 3,
			expect: nil,
		},
	}

	for _, test := range twoSumTest {
		get := removeElements(test.head, test.target)
		expect := test.expect
		if expect == nil {
			if get != nil {
				t.Fatalf("")
			} else {
				return
			}
		}
		for get.Next != nil || expect.Next != nil {
			if get.Val != expect.Val {
				t.Fatalf("expect:%v, get:%v", test.expect, get)
			}
			get = get.Next
			expect = expect.Next
		}
		if get.Next == nil && expect.Next != nil {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
		if get.Next != nil && expect.Next == nil {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
		if get.Val != expect.Val {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
