package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/symmetric-tree/description/

func isSymmetric01(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var traversal func(left *TreeNode, right *TreeNode) bool

	traversal = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		return traversal(left.Left, right.Right) && traversal(left.Right, right.Left)
	}

	return traversal(root.Left, root.Right)
}

func isSymmetric02(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		n1 := queue.Remove(queue.Front()).(*TreeNode)
		n2 := queue.Remove(queue.Front()).(*TreeNode)

		if n1 == nil && n2 == nil {
			continue // 继续比较
		} else if n1 == nil || n2 == nil {
			return false // 一定不是对称的
		}

		if n1.Val != n2.Val {
			return false
		}

		// 比较外侧
		queue.PushBack(n1.Left)
		queue.PushBack(n2.Right)

		// 比较里侧
		queue.PushBack(n1.Right)
		queue.PushBack(n2.Left)
	}
	return true
}

func TestIsSymmetric(t *testing.T) {
	case1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
	}
	case2 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 6}},
	}
	case3 := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}},
	}
	case4 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}
	var twoSumTest = []struct {
		name   string
		array  *TreeNode
		expect bool
	}{
		{name: "case1", array: case1, expect: false},
		{name: "case2", array: case2, expect: false},
		{name: "case3", array: case3, expect: false},
		{name: "case4", array: case4, expect: true},
		{name: "case5", array: nil, expect: true},
	}

	for _, test := range twoSumTest {
		get := isSymmetric02(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("%s expect:%v, get:%v", test.name, test.expect, get)
		}
	}
}
