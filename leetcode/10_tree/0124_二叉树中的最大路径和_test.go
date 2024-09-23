package _1_array

import (
	"math"
	"testing"
)

func maxPathSum(root *TreeNode) int {
	var traversal func(root *TreeNode) int

	mem := make(map[*TreeNode]int)
	traversal = func(root *TreeNode) int {
		if root == nil {
			mem[root] = math.MinInt32
			return math.MinInt32
		}

		rootVal := root.Val
		left, right, ok := 0, 0, false
		left, ok = mem[root.Left]
		if !ok {
			left = traversal(root.Left)
		}

		right, ok = mem[root.Right]
		if !ok {
			right = traversal(root.Right)
		}
		return max(rootVal, left, right, left+rootVal, right+rootVal, rootVal+left+right)
	}

	return traversal(root)
}

func TestMaxPathSum(t *testing.T) {
	var testdata = []struct {
		root *TreeNode
		want int
	}{
		{root: &TreeNode{Val: -3}, want: -3},
	}
	for _, tt := range testdata {
		get := maxPathSum(tt.root)
		if get != tt.want {
			t.Fatalf("want:%v, get:%v", tt.want, get)
		}
	}

}
