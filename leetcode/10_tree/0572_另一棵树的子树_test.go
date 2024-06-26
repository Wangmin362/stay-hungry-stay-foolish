package _1_array

// 地址：https://leetcode.cn/problems/subtree-of-another-tree/description/

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var traversal func(node *TreeNode)
	same := false
	traversal = func(node *TreeNode) {
		if same {
			return
		}
		if isSame(node, subRoot) {
			same = true
			return
		}

		if node.Left != nil {
			traversal(node.Left)
		}
		if node.Right != nil {
			traversal(node.Right)
		}
	}

	traversal(root)
	return same
}
