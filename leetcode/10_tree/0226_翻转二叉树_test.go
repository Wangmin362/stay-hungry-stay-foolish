package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/invert-binary-tree/description/

// 使用前序遍历，交换左右子节点——递归法
func invertTree(root *TreeNode) *TreeNode {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		traversal(node.Right)
		traversal(node.Left)
	}

	traversal(root)
	return root
}

// 前序遍历-迭代法
func invertTree02(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
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

// 后序遍历，递归法
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

// TODO 后序遍历-迭代法
func invertTree04(root *TreeNode) *TreeNode {
	return nil
}

func invertTree05(root *TreeNode) *TreeNode {
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
