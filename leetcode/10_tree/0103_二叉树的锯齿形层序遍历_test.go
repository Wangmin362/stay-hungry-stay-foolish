package _1_array

import (
	"container/list"
	"reflect"
	"slices"
	"testing"
)

// 地址：https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := list.New()
	queue.PushBack(root)
	l2r := true
	for queue.Len() > 0 {
		length := queue.Len()
		var temp []int
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if !l2r {
			slices.Reverse(temp)
		}
		res = append(res, temp)
		l2r = !l2r
	}
	return res
}

func TestZigzagLevelOrder(t *testing.T) {
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
		expect [][]int
	}{
		{array: case1, expect: [][]int{{4}, {9, 7}, {3, 2, 5, 6}}},
		{array: case2, expect: [][]int{{4}, {9, 7}, {3, 2, 6}}},
		{array: case3, expect: [][]int{{1}, {3}, {2}}},
		{array: nil, expect: nil},
	}

	for _, test := range twoSumTest {
		get := zigzagLevelOrder(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

	}
}
