package _1_array

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := &[]int{}

	InOrderTraversal(root, res)
	return *res
}

func InOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	InOrderTraversal(root.Right, res)
}

// TODO 迭代算法

func TestInOrderTraversal(t *testing.T) {
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

	var twoSumTest = []struct {
		array  *TreeNode
		expect []int
	}{
		{array: case1, expect: []int{3, 9, 2, 4, 5, 7, 6}},
		{array: case2, expect: []int{3, 9, 2, 4, 7, 6}},
		{array: case3, expect: []int{1, 2, 3}},
	}

	for _, test := range twoSumTest {
		get := inorderTraversal(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

		for i := 0; i < len(get); i++ {
			if get[i] != test.expect[i] {
				t.Fatalf("expect:%v, get:%v", test.expect, get)
			}
		}
	}
}
