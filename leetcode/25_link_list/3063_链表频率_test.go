package _1_array

func frequenciesOfElements(head *ListNode) *ListNode {
	c := make(map[int]int)
	for head != nil {
		c[head.Val]++
		head = head.Next
	}

	dummy := &ListNode{}
	curr := dummy
	for _, v := range c {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return dummy.Next
}
