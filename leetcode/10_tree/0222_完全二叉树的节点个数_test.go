package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/count-complete-tree-nodes/description/

// 后序遍历  递归
func countNodes01(root *TreeNode) int {
	var getNodes func(node *TreeNode) int

	getNodes = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lnodes := getNodes(node.Left)
		rnodes := getNodes(node.Right)
		return 1 + lnodes + rnodes
	}

	return getNodes(root)
}

// 迭代遍历  层序遍历
func countNodes02(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	nodes := 0
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nodes++
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return nodes
}
