package _1_array

import (
	"testing"
)

func postorderTraversal(root *TreeNode) []int {
	res := &[]int{}

	PostorderTraversal(root, res)
	return *res
}

func PostorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left, res)
	PostorderTraversal(root.Right, res)
	*res = append(*res, root.Val)
}

// TODO 迭代算法

func TestPostorderTraversal(t *testing.T) {
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
		{array: case1, expect: []int{3, 2, 9, 5, 6, 7, 4}},
		{array: case2, expect: []int{3, 2, 9, 6, 7, 4}},
		{array: case3, expect: []int{2, 3, 1}},
	}

	for _, test := range twoSumTest {
		get := postorderTraversal(test.array)
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
