package _1_array

import "container/list"

// 递归后续
func postOrderNTree01(root *Node) []int {
	var res []int
	var traversal func(node *Node)

	traversal = func(node *Node) {
		if node == nil {
			return
		}

		for _, n := range node.Children {
			traversal(n)
		}

		res = append(res, node.Val)
	}

	traversal(root)
	return res
}

// 后续遍历  迭代
func postOrderNTree(root *Node) []int {
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

			stack.PushBack(node)
			stack.PushBack(nil)

			for i := len(node.Children) - 1; i >= 0; i-- {
				if node.Children[i] != nil {
					stack.PushBack(node.Children[i])
				}
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*Node)
			res = append(res, node.Val)
		}
	}

	return res
}
