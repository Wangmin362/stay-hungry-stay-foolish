package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/invert-binary-tree/description/

// 前序遍历  递归
func invertTree01(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return root
}

// 中序遍历  递归
func invertTree02(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		node.Left, node.Right = node.Right, node.Left
		traversal(node.Left) // 这里必须是Left，因为中间节点进行了交换
	}

	traversal(root)
	return root
}

// 后续遍历  递归
func invertTree03(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}

	traversal(root)
	return root
}

// 前序迭代遍历
func invertTree04(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return root
}

// 中序迭代遍历  使用Nil标记法
func invertTree05(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back().Value
		if top != nil {
			node := stack.Remove(stack.Back()).(*TreeNode)

			if node.Right != nil {
				stack.PushBack(node.Right)
			}

			stack.PushBack(node)
			stack.PushBack(nil)

			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		} else {
			stack.Remove(stack.Back())
			node := stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

// 层序迭代遍历
func invertTree06(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return root
}

func TestInvertTree(t *testing.T) {
	case1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
	}
	case1expext := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
	}
	case2 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 6}},
	}
	case2expect := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
	}
	case3 := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}},
	}
	case3expext := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
	}

	var twoSumTest = []struct {
		array  *TreeNode
		expect *TreeNode
	}{
		{array: case1, expect: case1expext},
		{array: case2, expect: case2expect},
		{array: case3, expect: case3expext},
		{array: nil, expect: nil},
	}

	for _, test := range twoSumTest {
		get := invertTree05(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
