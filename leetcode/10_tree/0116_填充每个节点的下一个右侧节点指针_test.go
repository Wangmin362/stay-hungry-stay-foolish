package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var prev *Node
		for i := 0; i < length; i++ {
			top := queue.Remove(queue.Front()).(*Node)
			if prev != nil {
				prev.Next = top
			}
			prev = top
			if top.Left != nil {
				queue.PushBack(top.Left)
			}
			if top.Right != nil {
				queue.PushBack(top.Right)
			}
		}
	}
	return root
}
