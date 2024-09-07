package _3_tree

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/7298353c24cc42e3bd5f0e0bd3d1d759

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	node := t1
	node.Val += t2.Val
	node.Left = mergeTrees(t1.Left, t2.Left)
	node.Right = mergeTrees(t1.Right, t2.Right)
	return node
}

func TestMergeTrees(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t3 := mergeTrees(t1, t2)
	fmt.Println(t3)
}
