package _1_array

import "container/list"

// 地址：https://leetcode.cn/problems/n-ary-tree-preorder-traversal/description/

// 前序遍历N叉树，递归方法
func preorderNTree01(root *Node) []int {
	var res []int
	var traversal func(node *Node)

	traversal = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		for _, n := range node.Children {
			if n != nil {
				traversal(n)
			}
		}
	}

	traversal(root)
	return res
}

// 前序遍历，迭代法
func preOrderNTree02(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*Node)

		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			if node.Children[i] != nil {
				stack.PushBack(node.Children[i]) // 从后往前加，弹出来的时候才是从前往后
			}
		}
	}

	return res
}

func preOrderNTree03(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*Node)

			for i := len(node.Children) - 1; i >= 0; i-- {
				if node.Children[i] != nil {
					stack.PushBack(node.Children[i])
				}
			}

			stack.PushBack(node)
			stack.PushBack(nil)
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*Node)
			res = append(res, node.Val)
		}
	}
	return res
}
