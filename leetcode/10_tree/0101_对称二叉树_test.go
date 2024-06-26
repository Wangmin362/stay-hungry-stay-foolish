package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/symmetric-tree/description/

// 直接使用层序遍历的方式，使用队列实现层序遍历，使用数组来实现对称判断
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushFront(root)
	var node *TreeNode
	for queue.Len() > 0 {
		length := queue.Len() // 每一层的数量

		var stack []*TreeNode
		for i := 0; i < length; i++ {
			node = queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
				stack = append(stack, node.Left)
			} else {
				stack = append(stack, nil)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				stack = append(stack, node.Right)
			} else {
				stack = append(stack, nil)
			}
		}

		// 检查是否对称
		left, right := 0, len(stack)-1
		for left < right {
			l := stack[left]
			r := stack[right]
			if l == nil && r == nil {
				left++
				right--
				continue
			} else if l != nil && r != nil {
				if l.Val != r.Val {
					return false
				}
			} else {
				return false
			}

			left++
			right--
		}
	}

	return true
}

// 根本是对比左右子树是否时对称的
func isSymmetric01(root *TreeNode) bool {
	var compare func(left *TreeNode, right *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}

	if root == nil {
		return false
	}

	return compare(root.Left, root.Right)
}

// 迭代法
func isSymmetric02(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)
	for queue.Len() > 0 {
		n1 := queue.Remove(queue.Front()).(*TreeNode)
		n2 := queue.Remove(queue.Front()).(*TreeNode)
		if n1 == nil && n2 == nil {
			continue
		} else if n1 == nil && n2 != nil {
			return false
		} else if n1 != nil && n2 == nil {
			return false
		} else if n1.Val != n2.Val {
			return false
		}

		queue.PushBack(n1.Left)
		queue.PushBack(n2.Right)
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
		{name: "case5", array: nil, expect: false},
	}

	for _, test := range twoSumTest {
		get := isSymmetric02(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("%s expect:%v, get:%v", test.name, test.expect, get)
		}
	}
}
