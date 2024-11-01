package _1_array

func convertBiNode(root *TreeNode) *TreeNode {
	var traversal func(root *TreeNode)

	var prev *TreeNode
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Right)
		root.Right = prev
		prev = root
		traversal(root.Left)
		root.Left = nil
	}

	traversal(root)
	return prev
}
