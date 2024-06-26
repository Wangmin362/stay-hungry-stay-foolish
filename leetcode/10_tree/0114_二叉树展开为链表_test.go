package _1_array

// 地址：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/

func flatten(root *TreeNode) {
	var traversal func(node *TreeNode, prev *TreeNode)

	traversal = func(node *TreeNode, prev *TreeNode) {
		if node == nil {
			prev.Right = nil
			prev.Left = nil
			return
		}

		if prev != nil {
			prev.Right = node
			prev.Left = nil
		}

		traversal(node.Left, node)
		traversal(node.Right, node)
	}

	traversal(root, nil)
	return
}
