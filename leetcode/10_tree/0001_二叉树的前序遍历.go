package _1_array

import "container/list"

// 前序遍历  递归版本  中左右
func preOrderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 前序遍历 迭代版本  使用栈模拟迭代过程
func preOrderTraversal01(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode) // 由于是模拟栈，因此每次都需要从栈顶弹出一个元素

		res = append(res, node.Val) // 中  已经使用
		if node.Right != nil {
			stack.PushBack(node.Right) // 右 还未使用，下一次使用
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // 左 还未使用，下一次使用
		}
	}

	return res
}

// 前序遍历，迭代遍历  方式二，使用nil标记已经遍历过，但是还没有使用的元素
func preOrderTraversal02(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value // 看一下栈顶
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)

			if node.Right != nil {
				stack.PushBack(node.Right) // 右
			}

			if node.Left != nil {
				stack.PushBack(node.Left) // 左
			}

			stack.PushBack(node) // 中
			stack.PushBack(nil)  // 标记当前元素已经遍历过，但是还没有使用
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		}
	}

	return res
}
