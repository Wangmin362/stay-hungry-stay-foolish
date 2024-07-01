package _1_array

// 地址：https://leetcode.cn/problems/path-sum/description/

// 前序遍历  递归
func hasPathSum01(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var hasSum func(node *TreeNode, sum int) bool

	hasSum = func(node *TreeNode, sum int) bool {
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			return true
		}

		var lsum, rsum bool
		if node.Left != nil {
			lsum = hasSum(node.Left, sum)
			sum -= node.Val
		}
		if node.Right != nil {
			lsum = hasSum(node.Right, sum)
			sum -= node.Val
		}

		return lsum || rsum
	}

	return hasSum(root, 0)

}
func hasPathSum02(root *TreeNode, targetSum int) bool {
}
func hasPathSum03(root *TreeNode, targetSum int) bool {

}
