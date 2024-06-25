package _1_array

import (
	"container/list"
	"math"
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/

// 很简单，就是层序遍历
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		maxNum := math.MinInt
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val > maxNum {
				maxNum = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		res = append(res, maxNum)
	}

	return res
}

func TestLargestValues(t *testing.T) {
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
		expect []int
	}{
		{array: case1, expect: []int{4, 9, 6}},
		{array: case2, expect: []int{4, 9, 6}},
		{array: nil, expect: nil},
	}

	for _, test := range test {
		get := largestValues(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

	}
}
