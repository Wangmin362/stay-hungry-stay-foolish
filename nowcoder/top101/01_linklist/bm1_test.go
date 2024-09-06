package _1_linklist

// https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}
