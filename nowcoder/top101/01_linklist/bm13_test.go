package _1_linklist

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f

// 放入数组中对比，
func isPail00(head *ListNode) bool {
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left].Val != arr[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}

// 找到中间点，然后切分链表，反转后一半的链表，然后对比
func isPail(head *ListNode) bool {
	getLen := func(head *ListNode) int {
		var res int
		for head != nil {
			res++
			head = head.Next
		}
		return res
	}
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		for head != nil {
			nxt := head.Next
			head.Next = prev
			prev = head
			head = nxt
		}
		return prev
	}

	length := getLen(head)
	if length <= 1 {
		return true
	}
	mid := length / 2
	fast := head
	for mid > 1 {
		fast = fast.Next
		mid--
	}
	l1 := head
	l2 := fast.Next
	if length%2 == 1 {
		l2 = l2.Next
	}
	fast.Next = nil
	l2 = reverse(l2)
	for l1 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

func TestIsPail(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}}}}}}}
	//fmt.Println(isPail00(l))
	fmt.Println(isPail(l))
}
