package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/uncrossed-lines/description/

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// dp[i][j]定义为nums1[:i-1]和nums2[:j-1]公共子序列的最大长度
	// 若nums1[i-1]=nums2[j-1]那么，dp[i][j] = dp[i-1][j-1]+1
	// 否则需要看nums1[:i-2]和nums2[:j-1]的公共子序列以及nums1[:i-1][:j-2]的公共子序列
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	dp := make([][]int, len(nums1)+1)
	dp[0] = make([]int, len(nums2)+1)
	res := 0
	for i := 1; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}

	return res
}

func TestMaxUncrossedLines(t *testing.T) {
	fmt.Println(maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
}
