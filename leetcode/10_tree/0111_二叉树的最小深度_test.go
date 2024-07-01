package _1_array

import (
	"container/list"
	"math"
)

// 地址：https://leetcode.cn/problems/minimum-depth-of-binary-tree/

// 后序遍历  递归写法
func minDepth01(root *TreeNode) int {
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if node.Left == nil {
			return 1 + getDeepth(node.Right)
		} else if node.Right == nil {
			return 1 + getDeepth(node.Left)
		} else {
			ldeep := getDeepth(node.Left)
			rigth := getDeepth(node.Right)
			return 1 + int(math.Min(float64(ldeep), float64(rigth)))
		}
	}

	return getDeepth(root)
}

// 前序遍历 递归  回溯  TODO 前序遍历的精髓还没有掌握
func minDepth02(root *TreeNode) int {

	return 0
}

// 迭代  层序遍历
func minDepth03(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	deep := 0
	for queue.Len() > 0 {
		length := queue.Len()
		deep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return deep
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return deep
}
