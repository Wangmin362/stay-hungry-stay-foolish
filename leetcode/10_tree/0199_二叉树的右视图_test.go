package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

// 接替思路很简单，其实既是一个层序遍历，只不过每一层只要最后一个节点
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == length-1 { // 说明是最后一个节点
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res
}

func TestRightSideView(t *testing.T) {
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
		expect []int
	}{
		{array: case1, expect: []int{4, 7, 6}},
		{array: case2, expect: []int{4, 7, 6}},
		{array: case3, expect: []int{1, 3, 2}},
		{array: case4, expect: []int{1, 3, 4}},
		{array: case5, expect: []int{1, 3, 5}},
		{array: nil, expect: nil},
	}

	for _, test := range twoSumTest {
		get := rightSideView(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

	}
}
