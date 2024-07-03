package _1_array

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

// 后续遍历， 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var traversal func(node, p, q *TreeNode) *TreeNode

	traversal = func(node, p, q *TreeNode) *TreeNode {
		if node == q || node == p || node == nil {
			return node
		}

		left := traversal(node.Left, p, q)
		right := traversal(node.Right, p, q)
		if left != nil && right != nil {
			return node
		} else if left == nil {
			return right
		} else if right == nil {
			return left
		} else {
			return nil
		}
	}
	return traversal(root, p, q)
}
