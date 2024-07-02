package _1_array

// https://leetcode.cn/problems/merge-two-binary-trees/description/

// 合并后的树等于左子树和左子树合并，右子树和右子树合并，中间节点相加
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var traversal func(t1, t2 *TreeNode) *TreeNode

	traversal = func(t1, t2 *TreeNode) *TreeNode {
		if t1 == nil && t2 == nil {
			return nil
		} else if t1 == nil {
			return t2
		} else if t2 == nil {
			return t1
		}

		root := t1
		root.Val += t2.Val
		root.Left = traversal(t1.Left, t2.Left)
		root.Right = traversal(t1.Right, t2.Right)
		return root
	}

	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}

	node := root1
	node.Val += root2.Val
	node.Left = traversal(root1.Left, root2.Left)
	node.Right = traversal(root1.Right, root2.Right)
	return node
}

func mergeTrees02(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees02(root1.Left, root2.Left)
	root1.Right = mergeTrees02(root1.Right, root2.Right)
	return root1
}
