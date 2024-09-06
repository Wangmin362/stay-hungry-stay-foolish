package _1_linklist

// https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	getLen := func(head *ListNode) int {
		var res int
		for head != nil {
			head = head.Next
			res++
		}
		return res
	}
	length := getLen(pHead)
	if k > length || k <= 0 {
		return nil
	}

	dummy := &ListNode{Next: pHead}
	curr := dummy
	step := length - k
	for step > 0 {
		curr = curr.Next
		step--
	}

	return curr.Next
}
