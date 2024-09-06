package _1_linklist

// https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	reverseNodes := n - m + 1
	for m > 1 {
		p0 = p0.Next
		m--
	}

	var prev *ListNode
	curr := p0.Next
	for reverseNodes > 0 {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
		reverseNodes--
	}

	p0.Next.Next = curr
	p0.Next = prev
	return dummy.Next
}
