package _1_array

import (
	"container/list"
)

// 地址：https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

// 很简单，就是层序遍历
func maxDepth559(root *Node) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	deep := 0
	for queue.Len() > 0 {
		length := queue.Len()
		deep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for _, child := range node.Children {
				if child != nil {
					queue.PushBack(child)
				}
			}
		}
	}

	return deep
}
