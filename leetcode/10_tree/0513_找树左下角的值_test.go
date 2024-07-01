package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/sum-of-left-leaves/description/

// 前序遍历  递归
func findBottomLeftValue01(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDeep := 0
	val := 0
	var getDeepth func(node *TreeNode, deepth int)

	getDeepth = func(node *TreeNode, deepth int) {
		if node.Left == nil && node.Right == nil && deepth > maxDeep {
			maxDeep = deepth
			val = node.Val
			return
		}

		if node.Left != nil {
			deepth++
			getDeepth(node.Left, deepth)
			deepth--
		}
		if node.Right != nil {
			deepth++
			getDeepth(node.Right, deepth)
			deepth--
		}
	}

	getDeepth(root, 1)
	return val
}

// 层序遍历
func findBottomLeftValue03(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	val := 0
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if i == 0 {
				val = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return val
}
