package _1_array

// https://leetcode.cn/problems/convert-bst-to-greater-tree/description/

// 递归
func convertBST(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode)

	prev := 0
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Right)
		node.Val += prev
		prev = node.Val
		traversal(node.Left)
	}

	traversal(root)
	return root
}
