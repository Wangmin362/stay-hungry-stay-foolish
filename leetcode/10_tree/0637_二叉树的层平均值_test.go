package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

// 很简单，就是层序遍历
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var res []float64
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		sum := 0
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, float64(sum)/float64(length))
	}

	return res
}

func TestAverageOfLevels(t *testing.T) {
	case1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
	}
	case2 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 6}},
	}

	var test = []struct {
		array  *TreeNode
		expect []float64
	}{
		{array: case1, expect: []float64{4, 8, 4}},
		{array: case2, expect: []float64{4, 8, 3.6666666666666665}},
		{array: nil, expect: nil},
	}

	for _, test := range test {
		get := averageOfLevels(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

	}
}
