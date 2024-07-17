package main

import (
	"testing"
)

func rotate(nums []int, k int) {
}

func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	rotate([]int{1, 1, 1, 1, 2, 2, 2, 3}, 3)
}

type ListNode struct {
	Val  int
	Next *ListNode
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
