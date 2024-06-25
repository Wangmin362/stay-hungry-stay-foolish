package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var prev *Node
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if prev == nil {
				prev = node
			} else {
				prev.Next = node
				prev = node
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}
