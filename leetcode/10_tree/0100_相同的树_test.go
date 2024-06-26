package _1_array

// 地址：https://leetcode.cn/problems/same-tree/description/

// 只需要判断p的左子树和q的左子树是否相等。同时需要判断p的右子树和q的右子树是否相等
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 第二种解法，可以直接用迭代，同时跌倒两个树，使用层序遍历，然后判断每一层是否相等即可
