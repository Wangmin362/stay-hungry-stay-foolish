package _1_array

import "container/list"

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
func isSameTree02(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	queue := list.New()
	queue.PushBack(p)
	queue.PushBack(q)
	for queue.Len() > 0 {
		n1 := queue.Remove(queue.Front()).(*TreeNode)
		n2 := queue.Remove(queue.Front()).(*TreeNode)
		if n1 == nil && n2 == nil {
			continue
		} else if n1 == nil || n2 == nil {
			return false
		} else if n1.Val != n2.Val {
			return false
		}

		queue.PushBack(n1.Left)
		queue.PushBack(n2.Left)
		queue.PushBack(n1.Right)
		queue.PushBack(n2.Right)
	}
	return true
}
