package _3_tree

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/a69242b39baf45dea217815c7dedb52b

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	meetNil := false
	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		for _, node := range queue {
			if node == nil {
				meetNil = true
				continue
			}
			if meetNil { // 遇到nil之后，如果还遇到非空，说明不是一个完全二叉树
				return false
			}
			tmpQueue = append(tmpQueue, node.Left)
			tmpQueue = append(tmpQueue, node.Right)
		}
		queue = tmpQueue
	}

	return true
}

func TestIsCompleteTreeII(t *testing.T) {
	t1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	t3 := isCompleteTree(t1)
	fmt.Println(t3)
}
