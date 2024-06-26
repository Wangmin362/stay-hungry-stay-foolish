package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/count-complete-tree-nodes/description/

// 很简单，就是层序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	cnt := 0
	for queue.Len() > 0 {
		length := queue.Len()
		cnt += length
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

	return cnt
}
