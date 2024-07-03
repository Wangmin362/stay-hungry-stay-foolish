package _1_array

// https://leetcode.cn/problems/insert-into-a-binary-search-tree/

// 递归
func insertIntoBST01(root *TreeNode, val int) *TreeNode {
	var traversal func(node *TreeNode, val int) *TreeNode

	traversal = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{Val: val}
		}
		if node.Val > val {
			node.Left = traversal(node.Left, val)
		}
		if node.Val < val {
			node.Right = traversal(node.Right, val)
		}
		return node
	}

	return traversal(root, val)
}

// 迭代
func insertIntoBST02(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	curr := root
	for curr != nil {
		if curr.Val > val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				return root
			}
			curr = curr.Left
		}
		if curr.Val < val {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				return root
			}
			curr = curr.Right
		}
	}

	return root
}

// 迭代
func insertIntoBST03(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	curr := root
	prev := curr
	for curr != nil {
		prev = curr
		if curr.Val > val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if prev.Val > val {
		prev.Left = &TreeNode{Val: val}
	} else {
		prev.Right = &TreeNode{Val: val}
	}

	return root
}
