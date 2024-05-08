package _1_array

import (
	"testing"
)

// 图片地址：https://code-thinking.cdn.bcebos.com/gifs/102%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86.gif

func levelOrder(root *TreeNode) [][]int {
	return nil

}

func TestLevelOrder(t *testing.T) {
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
		{array: case1, expect: []int{3, 2, 9, 5, 6, 7, 4}},
		{array: case2, expect: []int{3, 2, 9, 6, 7, 4}},
		{array: case3, expect: []int{2, 3, 1}},
	}

	for _, test := range twoSumTest {
		get := levelOrder(test.array)
		if len(test.expect) != len(get) {
			t.Fatalf("expect:%v, get:%v", test.expect, get)
		}

		for i := 0; i < len(get); i++ {
			if get[i] != test.expect[i] {
				t.Fatalf("expect:%v, get:%v", test.expect, get)
			}
		}
	}
}
