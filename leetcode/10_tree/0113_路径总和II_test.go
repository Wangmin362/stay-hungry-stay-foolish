package _1_array

import (
	"reflect"
	"testing"
)

// 地址：https://leetcode.cn/problems/path-sum-ii/

// 前序遍历 递归
func pathSum01(root *TreeNode, targetSum int) [][]int {
	// golang中的slice是值拷贝，只要扩容了，底层数组就会发生变化，因此这里不能使用切片作为参数
	//var tarversal func(node *TreeNode, []int, sum int)
	if root == nil {
		return nil
	}

	var traversal func(node *TreeNode, sum int) // sum节点只包含当前节点之前节点的综合
	var path []int
	var res [][]int
	traversal = func(node *TreeNode, sum int) {
		sum += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == targetSum {
			tmp := make([]int, len(path))
			for idx, v := range path {
				tmp[idx] = v
			}
			res = append(res, tmp)
		}

		if node.Left != nil {
			traversal(node.Left, sum)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			traversal(node.Right, sum)
			path = path[:len(path)-1]
		}
	}

	traversal(root, 0)
	return res
}

func TestPathSum(t *testing.T) {
	case1 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}},
	}
	//case2 := &TreeNode{Val: 4,
	//	Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
	//	Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 6}},
	//}
	//case3 := &TreeNode{Val: 1,
	//	Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}},
	//}

	var twoSumTest = []struct {
		array  *TreeNode
		target int
		expect [][]int
	}{
		{array: case1, expect: [][]int{{4, 9, 3}, {4, 7, 5}}, target: 16},
		//{array: case2, expect: []int{3, 9, 2, 4, 7, 6}},
		//{array: case3, expect: []int{1, 2, 3}},
	}

	for _, test := range twoSumTest {
		get := pathSum01(test.array, test.target)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
