package _1_array

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	getLen := func(head *ListNode) int {
		var res int
		for head != nil {
			res++
			head = head.Next
		}
		return res
	}

	length := getLen(head)
	if length%2 == 1 && (length+1)/2 == k { // 说明需要交换的是中间节点，此时啥也不需要做
		return head
	}
	dummy := &ListNode{Next: head}
	curr := head
	step := k
	for step > 1 {
		curr = curr.Next
		step--
	}
	left := curr

	curr = head
	step = length - k
	for step > 1 {
		curr = curr.Next
		step--
	}
	right := curr
	if length>>1 < k {
		left, right = right, left // 说明
	}

	return dummy.Next
}
