package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/binary-tree-paths/description/

func binaryTreePaths(root *TreeNode) []string {
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
		get := binaryTreePaths(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
