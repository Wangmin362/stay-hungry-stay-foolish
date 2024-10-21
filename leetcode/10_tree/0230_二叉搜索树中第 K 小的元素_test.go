package _1_array

import "container/list"

// 迭代
func kthSmallest01(root *TreeNode, k int) int {
	var traversal func(root *TreeNode)
	var res int
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		traversal(root.Right)
	}
	traversal(root)
	return res
}

func kthSmallest02(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value // 看一下栈顶
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)

			if node.Right != nil {
				stack.PushBack(node.Right) // 右
			}

			stack.PushBack(node) // 中
			stack.PushBack(nil)  // 标记当前元素已经遍历过，但是还没有使用

			if node.Left != nil {
				stack.PushBack(node.Left) // 左
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			k--
			if k == 0 {
				return node.Val
			}
		}
	}

	return 0
}
