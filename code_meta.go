package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (r *ListNode) String() string {
	var res []string
	curr := r
	for curr != nil {
		res = append(res, strconv.Itoa(curr.Val))
		curr = curr.Next
	}
	return strings.Join(res, "->")
}

// TreeNode 二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node N叉树定义
type Node struct {
	Val      int
	Children []*Node
}
