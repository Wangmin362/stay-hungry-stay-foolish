package _3_tree

import (
	"container/list"
	"fmt"
	"testing"
)

// https://www.nowcoder.com/practice/a69242b39baf45dea217815c7dedb52b

// 递归
func isValidBST00(root *TreeNode) bool {
	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}

	traversal(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}

// 正宗递归
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var traversal func(root *TreeNode) bool
	var prev *TreeNode
	traversal = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := traversal(root.Left)
		if prev == nil {
			prev = root
		} else {
			if prev.Val > root.Val {
				return false
			}
			prev = root
		}

		right := traversal(root.Right)
		return left && right
	}
	return traversal(root)
}

// 迭代
func isValidBST03(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := list.New()
	stack.PushBack(root)
	var prev *TreeNode
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}

			stack.PushBack(node)
			stack.PushBack(nil)

			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			if prev == nil {
				prev = node
			} else {
				if prev.Val > node.Val {
					return false
				}
				prev = node
			}
		}

	}

	return true
}

func TestIsValidBST(t *testing.T) {
	t1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	t3 := isValidBST(t1)
	fmt.Println(t3)
}
