package _1_array

import (
	"math"
)

// 地址：https://leetcode.cn/problems/balanced-binary-tree/description/

// 后续遍历 递归
func isBalanced01(root *TreeNode) bool {
	// -1表示不是AVL树
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		ldeep := getDeepth(node.Left)
		if ldeep == -1 {
			return ldeep
		}
		rdeep := getDeepth(node.Right)
		if rdeep == -1 {
			return rdeep
		}
		if int(math.Abs(float64(ldeep-rdeep))) <= 1 {
			return 1 + int(math.Max(float64(ldeep), float64(rdeep)))
		} else {
			return -1 // 不是一颗AVL树
		}
	}

	if getDeepth(root) == -1 {
		return false
	}
	return true
}
