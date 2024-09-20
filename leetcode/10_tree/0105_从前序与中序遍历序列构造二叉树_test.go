package _1_array

import (
	"fmt"
	"testing"
)

// 地址：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(preorder []int, ps, pe int, inorder []int, is, ie int) *TreeNode

	build = func(preorder []int, ps, pe int, inorder []int, is, ie int) *TreeNode {
		if ps > pe || is > ie {
			return nil
		}
		root := &TreeNode{Val: preorder[ps]}
		idx := is // 找到中序遍历，根节点的索引，这样可以区分左子树和右子树的集合
		for idx <= ie {
			if inorder[idx] == preorder[ps] {
				break
			}
			idx++
		}
		length := idx - is // 长度不应该计算中间节点

		root.Left = build(preorder, ps+1, ps+1+length-1, inorder, is, idx-1)
		root.Right = build(preorder, ps+1+length-1+1, pe, inorder, idx+1, ie)
		return root
	}

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func TestName(t *testing.T) {
	tree := buildTree([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(tree)

	tree = buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree)
}
