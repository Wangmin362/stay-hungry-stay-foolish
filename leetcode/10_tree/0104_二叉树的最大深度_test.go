package _1_array

import (
	"container/list"
	"math"
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

// 递归 后序遍历
func maxDepth01(root *TreeNode) int {
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDeepth := getDeepth(node.Left)   // 左
		rightDeepth := getDeepth(node.Right) // 右
		if leftDeepth > rightDeepth {        // 中
			return 1 + leftDeepth
		} else {
			return 1 + rightDeepth
		}
	}

	return getDeepth(root)
}

// 递归  前序遍历
func maxDepth02(root *TreeNode) int {
	res := math.MinInt
	var getDeepth func(node *TreeNode, deepth int)

	getDeepth = func(node *TreeNode, deepth int) {
		if deepth > res {
			res = deepth
		}

		if node.Left != nil {
			deepth++
			getDeepth(node.Left, deepth)
			deepth--
		}
		if node.Right != nil {
			deepth++
			getDeepth(node.Right, deepth)
			deepth--
		}
	}

	if root == nil {
		return 0
	}

	getDeepth(root, 1)
	return res
}

// 迭代 层序遍历
func maxDepth03(root *TreeNode) int {
	deep := 0
	if root == nil {
		return deep
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		deep++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
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
		get := maxDepth03(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v, tree:%v", test.expect, get, test.array)
		}

	}
}
