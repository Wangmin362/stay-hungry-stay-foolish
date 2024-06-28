package _1_array

import (
	"container/list"
	"math"
)

// 地址：https://leetcode.cn/problems/sum-of-left-leaves/description/

// 解法一，使用层序遍历，每次记录每一次第一个值即可
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftVal int

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				leftVal = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return leftVal
}

// 解法二: 使用递归遍历  深度最深的左节点一定是要找的那个节点
func findBottomLeftValue01(root *TreeNode) int {
	deep := math.MinInt
	var res int
	var traversal func(node *TreeNode, dp int)
	traversal = func(node *TreeNode, dp int) {
		if node.Left == nil && node.Right == nil && dp > deep {
			deep = dp
			res = node.Val
		}

		if node.Left != nil {
			dp++
			traversal(node.Left, dp)
			dp-- // 回溯，左边遍历完成之后，就会退到上面的节点，此时就需要要减一
		}
		if node.Right != nil {
			dp++
			traversal(node.Right, dp)
			dp--
		}
	}

	traversal(root, 1)
	return res
}

// 解法二: 使用递归遍历  深度最深的左节点一定是要找的那个节点  精简回溯，应藏在参数当中
func findBottomLeftValue02(root *TreeNode) int {
	maxDeep := math.MinInt
	var res int
	var traversal func(node *TreeNode, deep int)
	traversal = func(node *TreeNode, deep int) {
		if node.Left == nil && node.Right == nil && deep > maxDeep {
			maxDeep = deep
			res = node.Val
		}
		if node.Left != nil {
			// 这里其实隐藏着回溯，应为这里的深度其实在遍历完成之后，外层的deep并没有真正的加一
			traversal(node.Left, deep+1)
		}
		if node.Right != nil {
			traversal(node.Right, deep+1)
		}
	}

	traversal(root, 1)
	return res
}
