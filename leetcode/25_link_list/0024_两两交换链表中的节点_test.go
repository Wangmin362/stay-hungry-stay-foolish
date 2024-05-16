package _1_array

import (
	"testing"
)

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	slow := dummy
	for slow.Next != nil && slow.Next.Next != nil {
		fast := slow.Next
		tmp := fast.Next
		fast.Next = slow
		slow.Next = tmp

		slow = slow.Next
	}

	return dummy.Next

}

func TestSwapPairs(t *testing.T) {
	var twoSumTest = []struct {
		head   *ListNode
		expect *ListNode
	}{
		{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
		{head: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
			expect: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{head: &ListNode{Val: 3},
			expect: &ListNode{Val: 3},
		},
	}

	for _, test := range twoSumTest {
		get := reverseList(test.head)
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
