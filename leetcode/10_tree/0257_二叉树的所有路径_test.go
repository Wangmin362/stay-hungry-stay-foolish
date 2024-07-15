package _1_array

import (
	"reflect"
	"strconv"
	"testing"
)

// 题目：https://leetcode.cn/problems/binary-tree-paths/description/

// 前序遍历 递归
func binaryTreePaths01(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var res []string
	var traversal func(node *TreeNode, path string)
	traversal = func(node *TreeNode, path string) {
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			return
		}

		path += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			traversal(node.Left, path)
		}
		if node.Right != nil {
			traversal(node.Right, path)
		}
	}

	traversal(root, "")
	return res
}

// 前序遍历 迭代 TODO 有待消化
func binaryTreePaths02(root *TreeNode) []string {
	return nil
}

func TestBinaryTreePaths(t *testing.T) {
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

	var test = []struct {
		array  *TreeNode
		expect *TreeNode
	}{
		{array: case1, expect: case1expext},
		{array: case2, expect: case2expect},
		{array: case3, expect: case3expext},
		{array: nil, expect: nil},
	}

	for _, test := range test {
		get := binaryTreePaths01(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
