package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/

// 后序遍历 递归
func maxDepth55901(root *Node) int {
	var getDeepth func(node *Node) int

	getDeepth = func(node *Node) int {
		if node == nil {
			return 0
		}

		maxDeep := 0
		for _, n := range node.Children {
			d := getDeepth(n)
			if d > maxDeep {
				maxDeep = d
			}
		}
		return 1 + maxDeep
	}

	return getDeepth(root)
}

// 前序遍历，递归，回溯
func maxDepth55902(root *Node) int {
	deep := 0
	if root == nil {
		return deep
	}
	var getDeepth func(node *Node, deepth int)

	getDeepth = func(node *Node, deepth int) {
		if deepth > deep {
			deep = deepth
		}

		deepth++
		for _, n := range node.Children {
			if n != nil {
				getDeepth(n, deepth)
			}
		}
		deepth--
	}

	getDeepth(root, 1)
	return deep
}

// 迭代 层序遍历
func maxDepth55903(root *Node) int {
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
			node := queue.Remove(queue.Front()).(*Node)
			for _, n := range node.Children {
				if n != nil {
					queue.PushBack(n)
				}
			}
		}
	}

	return deep
}
