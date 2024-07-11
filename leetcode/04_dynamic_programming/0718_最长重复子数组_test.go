package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/

func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]表四以i-1结尾的数组和j-1结尾的数组相同元素的最长长度
	// dp[i][j] = dp[i-1][j-1] + 1
	dp := make([][]int, len(nums1)+1)
	dp[0] = make([]int, len(nums2)+1)
	res := 0
	for i := 1; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}

func TestFindLength(t *testing.T) {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
