package _3_tree

import (
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/a9d0ecbacef9410ca97463e4a5c83be7

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	left := pRoot.Left
	right := pRoot.Right
	pRoot.Left = right
	pRoot.Right = left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}

func TestMirror(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t3 := Mirror(t1)
	fmt.Println(t3)
}
