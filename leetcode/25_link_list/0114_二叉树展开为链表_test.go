package _1_array

import (
	"container/list"
	"testing"
)

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/

// 先转为数组，然后设置右指针
func flatten01(root *TreeNode) {
	var tmp []*TreeNode
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp = append(tmp, node)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	for idx := 0; idx < len(tmp)-1; idx++ {
		tmp[idx].Left = nil
		tmp[idx].Right = tmp[idx+1]
	}

	if len(tmp) > 0 {
		root = tmp[0]
	}
	return
}

// 直接一次性搞定
func flatten02(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)
	var prev *TreeNode
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)

			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}

			stack.PushBack(node)
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			if prev != nil {
				prev.Left = nil
				prev.Right = node
			}
			prev = node
		}
	}
}

func TestFlatten(t *testing.T) {

}
