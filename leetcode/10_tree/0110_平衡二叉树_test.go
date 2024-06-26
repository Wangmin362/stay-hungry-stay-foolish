package _1_array

import (
	"container/list"
)

// 地址：https://leetcode.cn/problems/balanced-binary-tree/description/

// 很简单，两次层序遍历，第一次求出最小深度，第二次求出最大深度，相差不超过一就是AVL树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root)
	minDeep := 0
	for queue.Len() > 0 {
		length := queue.Len()
		minDeep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil || node.Right == nil {
				goto here
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

here:

	queue.PushBack(root)
	maxDeep := 0
	for queue.Len() > 0 {
		length := queue.Len()
		maxDeep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return maxDeep-minDeep <= 1
}
