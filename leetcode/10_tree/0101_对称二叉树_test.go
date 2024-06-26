package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/symmetric-tree/description/

// 直接使用层序遍历的方式，使用队列实现层序遍历，使用栈来实现对称判断
func isSymmetric(root *TreeNode) bool {

}

func TestIsSymmetric(t *testing.T) {
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
		expect bool
	}{
		{array: case1, expect: false},
		{array: case2, expect: false},
		{array: case3, expect: false},
		{array: nil, expect: true},
	}

	for _, test := range twoSumTest {
		get := isSymmetric(test.array)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}
	}
}
