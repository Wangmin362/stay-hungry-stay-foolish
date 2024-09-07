package _3_tree

import (
	"container/list"
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/8a2b2bf6c19b4f23a9bdb9b233eefa73

// 层序遍历
func maxDepth01(root *TreeNode) int {
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
			node := queue.Remove(queue.Front()).(*TreeNode)
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

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + max(left, right)
}

func TestMaxDepth(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(maxDepth(tree))
}
