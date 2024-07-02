package _1_array

import "container/list"

// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/

func searchBST(root *TreeNode, val int) *TreeNode {
	var tarversal func(node *TreeNode) *TreeNode
	tarversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == val {
			return node
		}

		if node.Left != nil && node.Val > val {
			return tarversal(node.Left)
		}
		if node.Right != nil && node.Val < val {
			return tarversal(node.Right)
		}
		return nil
	}

	return tarversal(root)
}

// 前序 迭代
func searchBST02(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Val == val {
			return node
		}
		if node.Val < val && node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Val > val && node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return nil
}

// 前序迭代, null指针标记法
func searchBST03(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Val < val && node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Val > val && node.Left != nil {
				stack.PushBack(node.Left)
			}

			stack.PushBack(node)
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Val == val {
				return node
			}
		}
	}

	return nil
}

func searchBST04(root *TreeNode, val int) *TreeNode {
	curr := root
	for curr != nil {
		if curr.Val == val {
			return curr
		}
		if curr.Val > val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}
