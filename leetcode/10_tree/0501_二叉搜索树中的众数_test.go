package _1_array

import "container/list"

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/

// 中序递归
func findMode01(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var traversal func(node *TreeNode)

	maxCnt := 0
	cnt := 0
	var prev *TreeNode
	var res []int
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if prev == nil {
			cnt = 1
		} else if prev.Val == node.Val {
			cnt++
		} else { // 说明是不同的元素
			cnt = 1
		}
		prev = node

		if cnt == maxCnt {
			res = append(res, node.Val)
		}
		if cnt > maxCnt {
			res = []int{node.Val}
			maxCnt = cnt
		}

		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 迭代  指针法
func findMode02(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	curr := root
	var prev *TreeNode
	cnt, maxCnt := 0, 0
	var res []int
	for curr != nil || stack.Len() > 0 {
		if curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			if prev == nil {
				cnt = 1
			} else if node.Val == prev.Val {
				cnt++
			} else {
				cnt = 1
			}
			prev = node
			if cnt == maxCnt {
				res = append(res, node.Val)
			}
			if cnt > maxCnt {
				res = []int{node.Val}
				maxCnt = cnt
			}
			curr = node.Right
		}
	}

	return res
}

// 迭代 nil标记法
func findMode03(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	stack.PushBack(root)
	cnt, maxCnt := 0, 0
	var prev *TreeNode
	var res []int
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
				cnt = 1
			} else if node.Val == prev.Val {
				cnt++
			} else {
				cnt = 1
			}
			prev = node

			if cnt == maxCnt {
				res = append(res, node.Val)
			}
			if cnt > maxCnt {
				maxCnt = cnt
				res = []int{node.Val}
			}
		}
	}
	return res
}
