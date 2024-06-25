package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deep := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		deep++
	}

	return deep
}

func TestMaxDepth(t *testing.T) {
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
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	case5 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}

	var twoSumTest = []struct {
		array  *TreeNode
		expect int
	}{
		{array: case1, expect: 3},
		{array: case2, expect: 3},
		{array: case3, expect: 3},
		{array: case4, expect: 3},
		{array: case5, expect: 3},
		{array: nil, expect: 0},
	}

	for _, test := range twoSumTest {
		get := maxDepth(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v, tree:%v", test.expect, get, test.array)
		}

	}
}
