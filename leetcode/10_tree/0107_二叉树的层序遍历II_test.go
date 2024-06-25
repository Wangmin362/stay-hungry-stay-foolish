package _1_array

import (
	"container/list"
	"reflect"
	"slices"
	"testing"
)

// 地址：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

// 使用层序遍历的方法，然后反转一下数组
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		temp := make([]int, length)
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			temp[i] = node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, temp)
	}

	slices.Reverse(res)
	return res
}

func TestLevelOrderBottom(t *testing.T) {
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
		{array: case1, expect: [][]int{{3, 2, 5, 6}, {9, 7}, {4}}},
		{array: case2, expect: [][]int{{3, 2, 6}, {9, 7}, {4}}},
		{array: case3, expect: [][]int{{2}, {3}, {1}}},
		{array: nil, expect: nil},
	}

	for _, test := range twoSumTest {
		get := levelOrderBottom(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

	}
}
