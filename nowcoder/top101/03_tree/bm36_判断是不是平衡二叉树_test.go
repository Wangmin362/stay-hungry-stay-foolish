package _3_tree

import (
	"container/list"
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222

func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(pRoot)
	maxDeep, minDeep := 0, 0
	findMinDeep := false
	for queue.Len() > 0 {
		length := queue.Len()
		maxDeep++
		if !findMinDeep {
			minDeep++
		}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				findMinDeep = true
			}
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

func TestIsBalanced_Solution(t *testing.T) {
	t1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	t3 := IsBalanced_Solution(t1)
	fmt.Println(t3)
}
