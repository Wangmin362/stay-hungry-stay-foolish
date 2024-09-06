package _1_linklist

// https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
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

	head1 = reverse(head1)
	head2 = reverse(head2)
	overflow := 0
	dummy := &ListNode{}
	curr := dummy
	for head1 != nil || head2 != nil {
		var v1, v2 int
		if head1 != nil {
			v1 = head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			v2 = head2.Val
			head2 = head2.Next
		}
		sum := v1 + v2 + overflow
		if sum > 9 {
			overflow = 1
			sum %= 10
		} else {
			overflow = 0
		}
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}
	if overflow == 1 {
		curr.Next = &ListNode{Val: 1}
	}

	return reverse(dummy.Next)
}
