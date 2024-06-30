package _1_array

import (
	"container/list"
	"testing"
)

// 一个方法搞定
func inorderTraversal02(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

// 迭代算法
// 还是需要使用栈的思想模拟中序遍历，核心思想就是需要先遍历到最左边的节点，直到左边节点为空了，就从栈中
// 去除这个元素，如果这个元素有右节点，那么依然还是要遍历到左边的节点，直到为空，重复这个动作
//        1
//     4     5
//   6   3  9

func inorderTraversal(root *TreeNode) []int {
	var res []int
	stack := list.New()
	curr := root
	for curr != nil || stack.Len() > 0 {
		if curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			curr = node.Right
		}
	}
	return res
}

func TestInOrderTraversal(t *testing.T) {
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
		expect []int
	}{
		{array: case1, expect: []int{3, 9, 2, 4, 5, 7, 6}},
		{array: case2, expect: []int{3, 9, 2, 4, 7, 6}},
		{array: case3, expect: []int{1, 2, 3}},
	}

	for _, test := range twoSumTest {
		get := inOrderTraversal02(test.array)

		for i := 0; i < len(get); i++ {
			if get[i] != test.expect[i] {
				t.Fatalf("expect:%v, get:%v", test.expect, get)
			}
		}
	}
}
