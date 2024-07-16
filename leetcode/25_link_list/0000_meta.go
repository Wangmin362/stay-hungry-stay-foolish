package _1_array

import "fmt"

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

func (r *ListNode) String() string {
	res := ""
	tmp := r
	for tmp != nil {
		res += fmt.Sprintf("%d -> ", tmp.Val)
		tmp = tmp.Next
	}
	res += "nil"
	return res
}

func linkListEqual(h1, h2 *ListNode) bool {
	h1Curr := h1
	h2Curr := h2
	for h1Curr != nil && h2Curr != nil {
		if h1Curr.Val != h2Curr.Val {
			return false
		}
		h1Curr = h1Curr.Next
		h2Curr = h2Curr.Next
	}
	if h1Curr == nil && h2Curr == nil {
		return true
	}
	return false
}
