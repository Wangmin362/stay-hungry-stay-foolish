package _1_array

import (
	"container/list"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/symmetric-tree/description/

// 直接使用层序遍历的方式，使用队列实现层序遍历，使用数组来实现对称判断
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.New()
	queue.PushFront(root)
	var node *TreeNode
	for queue.Len() > 0 {
		length := queue.Len() // 每一层的数量

		var stack []*TreeNode
		for i := 0; i < length; i++ {
			node = queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
				stack = append(stack, node.Left)
			} else {
				stack = append(stack, nil)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				stack = append(stack, node.Right)
			} else {
				stack = append(stack, nil)
			}
		}

		// 检查是否对称
		left, right := 0, len(stack)-1
		for left < right {
			l := stack[left]
			r := stack[right]
			if l == nil && r == nil {
				left++
				right--
				continue
			} else if l != nil && r != nil {
				if l.Val != r.Val {
					return false
				}
			} else {
				return false
			}

			left++
			right--
		}
	}

	return true
}

func TestIsSymmetric(t *testing.T) {
	//case1 := &TreeNode{Val: 4,
	//	Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
	//	Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
	//}
	//case2 := &TreeNode{Val: 4,
	//	Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
	//	Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 6}},
	//}
	//case3 := &TreeNode{Val: 1,
	//	Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}},
	//}
	case4 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}
	var twoSumTest = []struct {
		array  *TreeNode
		expect bool
	}{
		//{array: case1, expect: false},
		//{array: case2, expect: false},
		//{array: case3, expect: false},
		{array: case4, expect: true},
		{array: nil, expect: false},
	}

	for _, test := range twoSumTest {
		get := isSymmetric(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
