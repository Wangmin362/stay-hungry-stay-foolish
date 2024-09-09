package _0_basic

import "testing"

// https://leetcode.cn/problems/minimum-height-tree-lcci/description/?envType=problem-list-v2&envId=binary-search-tree&difficulty=EASY

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}

	var buildTree func(nums []int, start, end int) *TreeNode

	buildTree = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := start + (end-start)>>1
		root := &TreeNode{Val: nums[mid]}
		root.Left = buildTree(nums, start, mid-1)
		root.Right = buildTree(nums, mid+1, end)
		return root
	}

	return buildTree(nums, 0, len(nums)-1)
}

func TestSortedArrayToBST(t *testing.T) {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
