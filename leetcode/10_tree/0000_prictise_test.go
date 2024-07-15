package _1_array

import (
	"fmt"
	"testing"
)

func constructMaximumBinaryTree01(nums []int) *TreeNode {
	var build func(nums []int, begin, end int) *TreeNode

	build = func(nums []int, begin, end int) *TreeNode {
		if begin > end {
			return nil
		}
		maxIdx := begin
		for idx := begin; idx <= end; idx++ {
			if nums[idx] > nums[maxIdx] {
				maxIdx = idx
			}
		}

		root := &TreeNode{Val: nums[maxIdx]}
		root.Left = build(nums, begin, maxIdx-1)
		root.Right = build(nums, maxIdx+1, end)
		return root
	}

	return build(nums, 0, len(nums)-1)
}

func TestXxx(t *testing.T) {
	case1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}},
	}
	fmt.Println(case1)

	// root := buildTree01([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	root := constructMaximumBinaryTree01([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(root)
}
