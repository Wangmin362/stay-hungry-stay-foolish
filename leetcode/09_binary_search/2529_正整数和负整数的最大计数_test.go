package _9_binary_search

// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/

// 题目分析：题目需要计算负数的个数，以及正整数的个数，其中正整数不包含0，比较负数数量和正整数的数量，谁大返回谁的数量

// 思路一：直接以O(n)的方式遍历数组即可，分别统计负数的数量和正整数的数量
// 思路二：使用二分搜索，分别找到第一次小于0的位置（可以求出负数的数量），以及第一次大于0的位置（可以求出正整数的数量）

func maximumCount(nums []int) int {
	leftBorder := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}

	// 查找负数的数量，目标需要找到最后一次小于0的位置，其实就是找到第一次大于等于0的位置的左边
	negIdx := leftBorder(nums, 0) - 1
	negNum := negIdx + 1

	// 查找正整数的数量，目标就是要找到第一次大于等于1的位置
	posIdx := leftBorder(nums, 1)
	posNum := len(nums) - posIdx

	if negNum > posNum {
		return negNum
	}
	return posNum
}
