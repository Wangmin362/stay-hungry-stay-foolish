package _3_tree

import (
	"container/list"
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/5e2135f4d2b14eb8a5b06fab4c938635

// 递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// 遍历
func inorderTraversal01(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}

			stack.PushBack(node)
			stack.PushBack(nil)

			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}
	return res
}

func TestInorderTraversal(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(inorderTraversal01(tree))
}
