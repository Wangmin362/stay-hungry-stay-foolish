package _1_array

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/

// 中序递归
func findMode01(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var traversal func(node *TreeNode)

	maxCnt := -1
	var prev *TreeNode
	cnt := 1
	var res []int
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if prev != nil {
			if node.Val == prev.Val {
				cnt++
			} else {
				if maxCnt == -1 { // 修改初始状态
					maxCnt = cnt
				} else if cnt > maxCnt {
					maxCnt = cnt
					res = []int{node.Val}
				} else if cnt == maxCnt {
					res = append(res, node.Val)
				}
				cnt = 1
			}
		} else {
			cnt = 1
			res = []int{node.Val}
		}
		prev = node

		traversal(node.Right)
	}

	traversal(root)
	return res
}
