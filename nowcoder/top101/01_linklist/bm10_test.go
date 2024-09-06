package _1_linklist

// https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead1.Next == nil || pHead2 == nil || pHead2.Next == nil {
		return nil
	}
	getLen := func(head *ListNode) int {
		var res int
		for head != nil {
			head = head.Next
			res++
		}
		return res
	}

	p1Len := getLen(pHead1)
	p2Len := getLen(pHead2)

	if p1Len > p2Len { // 移动p1
		step := p1Len - p2Len
		for step > 0 {
			pHead1 = pHead1.Next
			step--
		}
	} else { // 移动p2
		step := p2Len - p1Len
		for step > 0 {
			pHead2 = pHead2.Next
			step--
		}
	}

	for pHead1 != nil && pHead2 != nil {
		if pHead1 == pHead2 {
			return pHead1
		}
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
	}

	return nil
}
