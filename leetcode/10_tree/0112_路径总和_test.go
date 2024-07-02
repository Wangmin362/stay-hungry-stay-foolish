package _1_array

import "container/list"

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

// 前序遍历 迭代 回溯，迭代想要带上回溯，就需要让迭代每个元素时，带上当前的情况。譬如这里时请求，就需要在栈中保存每个
// 节点当前的总和
func hasPathSum02(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	stack.PushBack(root)
	stack.PushBack(&TreeNode{Val: root.Val})
	for stack.Len() > 0 {
		val := stack.Remove(stack.Back()).(*TreeNode).Val
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Left == nil && node.Right == nil && val == targetSum {
			return true
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
			stack.PushBack(&TreeNode{Val: val + node.Right.Val})
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
			stack.PushBack(&TreeNode{Val: val + node.Left.Val})
		}
	}

	return false
}
