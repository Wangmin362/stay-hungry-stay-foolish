package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/maximum-binary-tree/description/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var traversal func(nums []int, begin, end int) *TreeNode

	traversal = func(nums []int, begin, end int) *TreeNode {
		if begin > end { // 边界不满足
			return nil
		}

		if begin == end { // 只有一个节点
			return &TreeNode{Val: nums[begin]}
		}

		maxIdx := begin
		for idx := begin; idx <= end; idx++ {
			if nums[idx] > nums[maxIdx] {
				maxIdx = idx
			}
		}
		node := &TreeNode{Val: nums[maxIdx]}

		node.Left = traversal(nums, begin, maxIdx-1)
		node.Right = traversal(nums, maxIdx+1, end)
		return node
	}

	return traversal(nums, 0, len(nums)-1)
}
func TestConstructMaximumBinaryTree(t *testing.T) {
	tree := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(tree)
}
