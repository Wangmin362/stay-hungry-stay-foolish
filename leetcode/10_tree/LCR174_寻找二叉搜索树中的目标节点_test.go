package _1_array

// 二叉搜索树从小到大遍历顺序为：左中右， 从大到小遍历顺序为：右中左
func findTargetNodeDfs(root *TreeNode, cnt int) int {
	var traversal func(root *TreeNode)

	var res int
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Right)
		cnt--
		if cnt == 0 {
			res = root.Val
		}
		traversal(root.Left)
	}
	traversal(root)
	return res
}

// 迭代
func findTargetNodeRecursive(root *TreeNode, cnt int) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0, 64)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top != nil {
			no := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if no.Left != nil {
				stack = append(stack, no.Left)
			}

			stack = append(stack, no)
			stack = append(stack, nil)

			if no.Right != nil {
				stack = append(stack, no.Right)
			}
		} else {
			no := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			cnt--
			if cnt == 0 {
				return no.Val
			}
		}
	}
	return -1
}
