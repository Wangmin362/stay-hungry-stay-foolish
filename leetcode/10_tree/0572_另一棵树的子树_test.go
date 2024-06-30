package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/subtree-of-another-tree/description/

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}

// 递归
func isSubtree01(root *TreeNode, subRoot *TreeNode) bool {
	same := false
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if same {
			return
		}
		if node == nil {
			return
		}

		if isSame(node, subRoot) {
			same = true
			return
		}

		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return same
}

// 迭代
func isSubtree02(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		if subRoot == nil {
			return true
		} else {
			return false
		}
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		if isSame(node, subRoot) {
			return true
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return false
}
