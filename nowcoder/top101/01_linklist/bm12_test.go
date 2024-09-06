package _1_linklist

import (
	"sort"
	"testing"
)

// https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08

// 转为数组，然后排序
func sortInList(head *ListNode) *ListNode {
	var arr []*ListNode
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

func TestSortInList(t *testing.T) {
	l := &ListNode{Val: 1, Next: &ListNode{Val: -2, Next: &ListNode{Val: 9, Next: &ListNode{Val: 3}}}}
	sortInList(l)

}
