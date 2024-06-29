package _1_array

import "container/list"

// 二叉树的前序遍历 递归版本
func preOrderTraversal01(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 二叉树的后续遍历，递归版本
func inorderTraversal01(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 二叉树的后续遍历 递归版本
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

// 二叉树的前序遍历  迭代版本
func preOrderTraversalIter01(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return res
}

func preOrderTraversal02(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode) // 从栈里弹出来

			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			stack.PushBack(node) // 从栈里弹出来的顺序为中左右
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back()) // 弹出nil
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}

	return res
}
