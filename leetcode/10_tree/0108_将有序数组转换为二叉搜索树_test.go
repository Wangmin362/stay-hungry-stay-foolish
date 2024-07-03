package _1_array

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/

// 思路：选取中间节点作为根节点，左边数组构造左子树，右边数组构造右子树即可
func sortedArrayToBST(nums []int) *TreeNode {
	var traversal func(nums []int, left, right int) *TreeNode

	traversal = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)>>1
		root := &TreeNode{Val: nums[mid]}
		root.Left = traversal(nums, left, mid-1)
		root.Right = traversal(nums, mid+1, right)
		return root
	}

	return traversal(nums, 0, len(nums)-1)
}
