package _1_array

import (
	"testing"
)

func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}

	PreorderTraversal(root, res)
	return *res
}

func PreorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	PreorderTraversal(root.Left, res)
	PreorderTraversal(root.Right, res)
}

// TODO 迭代算法

func TestPreorderTraversal(t *testing.T) {
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
		{array: case1, expect: []int{4, 9, 3, 2, 7, 5, 6}},
		{array: case2, expect: []int{4, 9, 3, 2, 7, 6}},
		{array: case3, expect: []int{1, 3, 2}},
	}

	for _, test := range twoSumTest {
		get := preorderTraversal(test.array)
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
