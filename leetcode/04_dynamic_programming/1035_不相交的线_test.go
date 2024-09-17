package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/uncrossed-lines/description/

// 题目分析：本质上及时再求两个数组的最长公共子序列
// 明确定义：dp[i][j]表示nums1[0:i-1]和nums2[0:j-1]的最长公共子序列
// 递推公式：if nums1[i-1]==nums2[j-2] dp[i][j] = dp[i-1][j-1]+1 else dp[i][j] = max(dp[i][j-1], dp[i-1][j])
// 初始化：dp[0][j] = 0, dp[j][0] = 0
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	var res int
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}

func TestMaxUncrossedLines(t *testing.T) {
	var testdata = []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{nums1: []int{1, 4, 2}, nums2: []int{1, 2, 4}, want: 2},
		{nums1: []int{3, 3, 2}, nums2: []int{3, 3, 1, 2}, want: 3},
	}
	for _, tt := range testdata {
		get := maxUncrossedLines(tt.nums1, tt.nums2)
		if get != tt.want {
			t.Fatalf("nums1:%v, nums2:%v, want%v, get:%v", tt.nums1, tt.nums2, tt.want, get)
		}
	}
}
