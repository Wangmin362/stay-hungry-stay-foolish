package _1_linklist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var res []string
	curr := n
	for curr != nil {
		res = append(res, strconv.Itoa(curr.Val))
		curr = curr.Next
	}
	return strings.Join(res, "->")
}
