package _1_linklist

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	evenFirst := even
	curr := even.Next
	flag := true
	for curr != nil {
		if flag {
			odd.Next = curr
			odd = odd.Next
		} else {
			even.Next = curr
			even = even.Next
		}
		curr = curr.Next
		flag = !flag
	}
	even.Next = nil
	odd.Next = evenFirst

	return head
}

func TestOddEvenList(t *testing.T) {
	l := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5}}}}}
	fmt.Println(oddEvenList(l))
}
