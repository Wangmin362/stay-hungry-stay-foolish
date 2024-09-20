package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

func buildTree106(inorder []int, postorder []int) *TreeNode {
	var build func(inorder []int, is, ie int, postorder []int, ps, pe int) *TreeNode

	build = func(inorder []int, is, ie int, postorder []int, ps, pe int) *TreeNode {
		if is > ie || ps > ie {
			return nil
		}

		root := &TreeNode{Val: postorder[pe]}
		idx := is // 找到中序遍历根节点的索引号，拆分左子树和右子树的集合
		for idx <= ie {
			if inorder[idx] == postorder[pe] {
				break
			}
			idx++
		}
		length := idx - is // 左子树长度，不包含根节点，所以无需减一

		root.Left = build(inorder, is, idx-1, postorder, ps, ps+length-1)
		root.Right = build(inorder, idx+1, ie, postorder, ps+length, pe-1)
		return root
	}

	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func TestName11(t *testing.T) {
	tree := buildTree106([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(tree)
}
