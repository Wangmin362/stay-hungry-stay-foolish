package _1_array

// https://leetcode.cn/problems/trim-a-binary-search-tree/description/

// 思路：按照删除BST节点的方式来删除
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var traversal func(node *TreeNode, low, high int) *TreeNode

	traversal = func(node *TreeNode, low, high int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < low { // 边界为[low, high]
			return traversal(node.Right, low, high)
		} else if node.Val > high {
			return traversal(node.Left, low, high)
		}

		node.Left = traversal(node.Left, low, high)
		node.Right = traversal(node.Right, low, high)
		return node
	}

	return traversal(root, low, high)
}
