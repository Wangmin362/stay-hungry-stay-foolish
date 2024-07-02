package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func buildTree02(inorder []int, postorder []int) *TreeNode {
	var traversal func(inorder []int, inBegin, inEnd int, postorder []int, postBegin, postEnd int) *TreeNode

	traversal = func(inorder []int, inBegin, inEnd int, postorder []int, postBegin, postEnd int) *TreeNode {
		if inBegin > inEnd || postBegin > postEnd {
			return nil
		}

		// 中序 左中右 9,3,15,20,7
		// 后续 左右中 9,15,7,20,3
		node := &TreeNode{Val: postorder[postEnd]}
		if inBegin == inEnd { // 只有一个节点
			return node
		}

		// 切分中序遍历左右子树
		idx := inBegin
		for ; idx <= inEnd; idx++ {
			if inorder[idx] == postorder[postEnd] {
				break
			}
		}
		length := idx - inBegin // 左子树的长度
		limitor := postBegin + length - 1
		node.Left = traversal(inorder, inBegin, idx-1, postorder, postBegin, limitor)
		node.Right = traversal(inorder, idx+1, inEnd, postorder, limitor+1, postEnd-1)
		return node
	}

	return traversal(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func TestName11(t *testing.T) {
	tree := buildTree02([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(tree)
}
