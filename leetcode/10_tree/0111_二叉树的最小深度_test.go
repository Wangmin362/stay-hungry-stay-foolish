package _1_array

import (
	"container/list"
)

// 地址：https://leetcode.cn/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deep := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		deep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return deep
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return deep
}
