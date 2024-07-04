package _1_array

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

// 后序遍历  虽然利用了二叉搜索树的性质，但是没有完全利用性质
func lowestCommonAncestorBST01(root, p, q *TreeNode) *TreeNode {
	var traversal func(node, p, q *TreeNode) *TreeNode

	traversal = func(node, p, q *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}

		var left, right *TreeNode
		if node.Val < p.Val && node.Val < q.Val { // 不需要向下查找

		} else {
			left = traversal(node.Left, p, q)
		}

		if node.Val > p.Val && node.Val > q.Val {

		} else {
			right = traversal(node.Right, p, q)
		}
		if left != nil && right != nil {
			return node
		} else if left == nil {
			return right
		} else if right == nil {
			return left
		} else {
			return nil
		}
	}

	return traversal(root, p, q)
}

// 递归 前序遍历
func lowestCommonAncestorBST02(root, p, q *TreeNode) *TreeNode {
	var traversal func(node, p, q *TreeNode) *TreeNode

	traversal = func(node, p, q *TreeNode) *TreeNode {
		if node == nil || (node.Val >= p.Val && node.Val <= q.Val) || (node.Val >= q.Val && node.Val <= p.Val) {
			return node
		}

		if node.Val > p.Val && node.Val > q.Val {
			left := traversal(node.Left, p, q)
			if left != nil {
				return left
			}
		}
		if node.Val < p.Val && node.Val < q.Val {
			right := traversal(node.Right, p, q)
			if right != nil {
				return right
			}
		}
		return nil
	}

	return traversal(root, p, q)
}

// 迭代
func lowestCommonAncestorBST03(root, p, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if (curr.Val >= p.Val && curr.Val <= q.Val) || (curr.Val >= q.Val && curr.Val <= p.Val) { // curr.Val在[p, q]或者 [q, p]之间
			return curr
		}
		if curr.Val < p.Val && curr.Val < q.Val { // curr.Val在[p, q]或者[q, p]左边
			curr = curr.Right
		}
		if curr.Val > p.Val && curr.Val > q.Val { // curr.Val在[p, q]或者[q, p]右边
			curr = curr.Left
		}
	}

	return nil
}
