package _1_array

// 地址：https://leetcode.cn/problems/sum-of-left-leaves/description/

func sumOfLeftLeaves(root *TreeNode) int {
	var traversal func(node *TreeNode, isLeft bool)
	sum := 0
	traversal = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
		}
		traversal(node.Left, true)
		traversal(node.Right, false)
	}

	traversal(root, false)
	return sum
}
