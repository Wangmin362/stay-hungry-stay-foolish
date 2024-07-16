package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/odd-even-linked-list/description/

func oddEvenList(head *ListNode) *ListNode {
	even, odd := &ListNode{}, &ListNode{}

	curr := head
	cnt := 0
	evenCur, oddCur := even, odd
	for curr != nil {
		cnt++
		if cnt%2 == 1 {
			evenCur.Next = curr
			evenCur = evenCur.Next
		} else {
			oddCur.Next = curr
			oddCur = oddCur.Next
		}
		curr = curr.Next
	}
	evenCur.Next = odd.Next
	oddCur.Next = nil
	return even.Next
}

func TestOddEvenList(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	oddEvenList(head)
	fmt.Println(head)
}
