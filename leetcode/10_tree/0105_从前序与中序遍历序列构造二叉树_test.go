package _1_array

import (
	"fmt"
	"testing"
)

// 地址：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {

	var traversal func(preorder []int, preBegin, preEnd int, inorder []int, inBegin, inEnd int) *TreeNode

	// 前序遍历：中左右  3,9,20,15,7
	// 中序遍历：左中右  9,3,15,20,7
	traversal = func(preorder []int, preBegin, preEnd int, inorder []int, inBegin, inEnd int) *TreeNode {
		if preBegin > preEnd || inBegin > inEnd {
			return nil
		}
		node := &TreeNode{Val: preorder[preBegin]}
		if preBegin == preEnd { // 如果只有一个节点，说明是叶子节点，直接返回
			return node
		}

		// 切分中序遍历，找到左子树和右子树集合
		idx := 0
		for idx = inBegin; idx <= inEnd; idx++ {
			if inorder[idx] == preorder[preBegin] {
				break // 在中序遍历中已经找到了中序节点
			}
		}
		length := idx - inBegin // 计算左子树一共有多少元素
		preLimitor := preBegin + length

		node.Left = traversal(preorder, preBegin+1, preLimitor, inorder, inBegin, idx-1) // taversal(前序遍历左子树，后序遍历左子树)
		node.Right = traversal(preorder, preLimitor+1, preEnd, inorder, idx+1, inEnd)    // traversal(前序遍历右子树， 后序遍历右子树)
		return node

	}

	return traversal(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func TestName(t *testing.T) {
	tree := buildTree([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(tree)
}
