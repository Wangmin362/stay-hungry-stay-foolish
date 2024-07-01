package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/sum-of-left-leaves/description/

// 前序遍历  递归
func sumOfLeftLeaves01(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var traversal func(node *TreeNode, isLeft bool)

	sum := 0
	traversal = func(node *TreeNode, isLeft bool) {
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
			return
		}

		if node.Left != nil {
			traversal(node.Left, true)
		}
		if node.Right != nil {
			traversal(node.Right, false)
		}
	}

	traversal(root, false)
	return sum
}

// 前序遍历  迭代
func sumOfLeftLeaves02(root *TreeNode) int {
	sum := 0
	stack := list.New()
	stack.PushBack(root)
	right := &TreeNode{}
	stack.PushBack(right)
	for stack.Len() > 0 {
		isLeft := stack.Remove(stack.Back()) == nil
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
			stack.PushBack(right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
			stack.PushBack(nil)
		}
	}

	return sum
}

// 层序遍历
func sumOfLeftLeaves03(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	right := &TreeNode{}
	queue.PushBack(right)
	sum := 0
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		isLeft := queue.Remove(queue.Front()) == nil
		if node.Left == nil && node.Right == nil && isLeft {
			sum += node.Val
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
			queue.PushBack(nil)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			queue.PushBack(right)
		}
	}

	return sum
}
