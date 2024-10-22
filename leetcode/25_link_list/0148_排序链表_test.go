package _1_array

import "sort"

// 直接使用数组排序，空间复杂度为O（n）
func sortList01(head *ListNode) *ListNode {
	arr := make([]*ListNode, 0, 64)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})

	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = arr[i+1]
	}

	return arr[0]
}

// 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 如果为空，或者就一个节点直接返回，不需要排序
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil { // 找到链表中间节点，如果是偶数，则为中间左边节点，如果是奇数则为中间节点
		slow = slow.Next
		fast = fast.Next.Next
	}

	h1, h2 := head, slow.Next
	slow.Next = nil                     // 断开链
	h1, h2 = sortList(h1), sortList(h2) // 继续使用归并排序
	// 到了这一步，我们拿到的就是排好序的两个链表，此时需要合并两个有序链表
	dummy := &ListNode{}
	curr := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			curr.Next = h1
			h1 = h1.Next
		} else {
			curr.Next = h2
			h2 = h2.Next
		}
		curr = curr.Next
	}
	if h1 != nil {
		curr.Next = h1
	}
	if h2 != nil {
		curr.Next = h2
	}

	return dummy.Next
}
