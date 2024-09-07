package _3_tree

import (
	"container/list"
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/5e2135f4d2b14eb8a5b06fab4c938635

// 递归
func preorderTraversal00(root *TreeNode) []int {
	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

// 栈遍历
func preorderTraversal02(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, top.Val)
		if top.Right != nil {
			stack.PushBack(top.Right)
		}
		if top.Left != nil {
			stack.PushBack(top.Left)
		}
	}
	return res
}

// 迭代统一写法  使用Null标记法
func preorderTraversal03(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}

			stack.PushBack(node)
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}

	return res
}

func TestPreOrderTraversal(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(preorderTraversal03(tree))
}
