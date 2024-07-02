package _1_array

import (
	"container/list"
	"math"
)

// https://leetcode.cn/problems/validate-binary-search-tree/description/

// 使用中序遍历 结果保存为数组，只要数组时有序的就是正确的
func isValidBST01(root *TreeNode) bool {
	var traversal func(node *TreeNode)

	var arr []int
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		arr = append(arr, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	for idx := range arr {
		if idx == len(arr)-1 {
			break
		}

		if arr[idx] >= arr[idx+1] {
			return false
		}
	}

	return true
}

// 保证当前节点的左子树所有节点小于当前节点，右子树所有节点大于当前节点
func isValidBST02(root *TreeNode) bool {
	var isValid func(node *TreeNode) bool
	var compare func(node *TreeNode, val int, isLessThan bool) bool // 用于完成当前二叉树是否全部小于目标值，或者全部大于目标值

	compare = func(node *TreeNode, val int, isLessThan bool) bool {
		if node == nil {
			return true
		}
		if isLessThan { // 要求二叉树的全部节点小于目标值
			if node.Val < val {
				lres := compare(node.Left, val, isLessThan)
				rres := compare(node.Right, val, isLessThan)
				return lres && rres
			} else {
				return false
			}
		} else { // 要求二叉树的全部节点大于目标值
			if node.Val > val {
				lres := compare(node.Left, val, isLessThan)
				rres := compare(node.Right, val, isLessThan)
				return lres && rres
			} else {
				return false
			}
		}
	}

	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !compare(node.Left, node.Val, true) || !compare(node.Right, node.Val, false) {
			return false
		}
		return isValid(node.Left) && isValid(node.Right)
	}

	return isValid(root)
}

// 使用中序遍历  中序遍历时节点一定是递增状态
func isValidBST03(root *TreeNode) bool {
	minVal := math.MinInt
	var isValid func(node *TreeNode) bool

	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftValid := isValid(node.Left)
		if node.Val > minVal {
			minVal = node.Val
		} else {
			return false
		}
		rightValid := isValid(node.Right)
		return leftValid && rightValid
	}

	return isValid(root)
}

// 继续优化
func isValidBST04(root *TreeNode) bool {
	var minVal *TreeNode
	var isValid func(node *TreeNode) bool

	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftValid := isValid(node.Left)
		if minVal == nil {
			minVal = node
		} else if node.Val > minVal.Val {
			minVal = node
		} else {
			return false
		}
		rightValid := isValid(node.Right)
		return leftValid && rightValid
	}

	return isValid(root)
}

// 迭代中序遍历
func isValidBST05(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var prev *TreeNode
	stack := list.New()
	stack.PushBack(root)
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
			if prev != nil && node.Val <= prev.Val {
				return false
			}
			prev = node // 更新最小值
		}
	}

	return true
}
