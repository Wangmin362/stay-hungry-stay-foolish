package _1_array

import "container/list"

// 中序遍历，递归版本
func inOrderTraversal01(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 中序遍历  迭代版本,使用指针遍历树
func inOrderTraversal02(root *TreeNode) []int {
	var res []int
	stack := list.New()
	curr := root
	for curr != nil || stack.Len() > 0 {
		if curr != nil { // 只要不为空，就一直向左遍历，直到遍历到最左边的节点，因为中序遍历时左中右
			stack.PushBack(curr)
			curr = curr.Left
		} else {
			curr = stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, curr.Val)
			curr = curr.Right
		}
	}

	return res
}

// 中序遍历，迭代方式。  使用nil标记法，使用nil来标记已经访问过但是还没有使用过的元素
func inOrderTraversal03(root *TreeNode) []int {
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
				stack.PushBack(node.Right) // 右
			}

			stack.PushBack(node) // 中
			stack.PushBack(nil)

			if node.Left != nil {
				stack.PushBack(node.Left) // 左
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}

	return res
}
