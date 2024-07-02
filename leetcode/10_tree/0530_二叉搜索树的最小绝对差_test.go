package _1_array

import (
	"container/list"
	"math"
)

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/

// 中序遍历  递归  两个指针，一个前一个节点，一个指向当前节点
func getMinimumDifference01(root *TreeNode) int {
	var traversal func(node *TreeNode)

	var prev *TreeNode // 指向前一个节点
	res := math.MaxInt
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if prev != nil && (node.Val-prev.Val) < res {
			res = node.Val - prev.Val
		}
		prev = node
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// 中序遍历 递归 两个指针
func getMinimumDifference02(root *TreeNode) int {
	stack := list.New()
	stack.PushBack(root)
	var prev *TreeNode
	res := math.MaxInt
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
			if prev != nil && (node.Val-prev.Val) < res {
				res = node.Val - prev.Val
			}
			prev = node
		}
	}

	return res
}
