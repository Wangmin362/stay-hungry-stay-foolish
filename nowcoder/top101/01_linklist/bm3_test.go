package _1_linklist

// https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			curr.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			curr.Next = pHead1
			pHead1 = pHead1.Next
		}

		curr = curr.Next
	}
	if pHead1 != nil {
		curr.Next = pHead1
	}
	if pHead2 != nil {
		curr.Next = pHead2
	}

	return dummy.Next
}
