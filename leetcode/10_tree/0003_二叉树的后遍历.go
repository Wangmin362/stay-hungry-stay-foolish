package _1_array

import (
	"container/list"
	"slices"
)

// 后序遍历，递归方式
func postOrderTraversal01(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}

	traversal(root)
	return res
}

// 后序遍历 迭代,把前序遍历的迭代稍微修改一下，变为中右左，反转之后就变为左右中
func postOrderTraversal02(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}

	slices.Reverse(res)
	return res
}

// 后序遍历 迭代方式。 使用Nil标记法，用于标记已经访问过，但是没有使用的元素
func postTraversal03(root *TreeNode) []int {
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

			stack.PushBack(node) // 中
			stack.PushBack(nil)

			if node.Right != nil {
				stack.PushBack(node.Right) // 右
			}

			if node.Left != nil {
				stack.PushBack(node.Left) // 左
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}

	return res
}
