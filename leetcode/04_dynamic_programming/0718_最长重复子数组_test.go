package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/

// 解法一：暴露搜索
// 解法二：滑动窗口
// 解法三：动态规划

func findLength01(nums1 []int, nums2 []int) int {
	var res int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			ii, jj, cnt := i, j, 0
			for ii < len(nums1) && jj < len(nums2) && nums1[ii] == nums2[jj] {
				ii++
				jj++
				cnt++
				res = max(res, cnt)
			}
		}
	}
	return res
}

// 题目分析：要求求出nums1和nums2的重复子数组的最长长度，由于是子数组，那么数组中的值一定是连续的
// 明确定义：dp[i][j]表示以i-1为结尾的数组nums1和以j-1为结尾的子数组nums2的最长子数组的长度，也就是nums1[0:i-1]和nums2[0:j-1]子数组的
// 最长长度
// 递推公式：dp[i][j] = dp[i-1][j-1] + 1, 显然，只有当nums1[i-1] = nums2[j-1]的时候，这个公式才成立。根据定义，dp[i-1][j-1]表示的是
// 以i-2结尾的子数组nums1[0:i-2]和以j-2为结尾的子数组nums2[0:j-2]的最长长度，当nums1[i-1] = nums2[j-1]时，显然dp[i][j] = dp[i-1][j-1] + 1
// 初始化：根据递推公式，由于dp[i][j]表示的时nums1[0:i-1]和nums2[0:j-1]的最长长度，并且dp[i][j] = dp[i-1][j-1] + 1, 所以i,j都必须从
// 1开始，dp[0][j]和dp[i][0]都是没有意义的，初始化为0就可以
// dp数组大小：[len(nums1)][len(nums2)]
// 返回值：过程中取最大值
func findLength03(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	var res int
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(res, dp[i][j])
		}

	}

	return res
}

func TestFindLength(t *testing.T) {
	var testdata = []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{nums1: []int{1, 2, 3, 2, 1}, nums2: []int{3, 2, 1, 4, 7}, want: 3},
	}
	for _, tt := range testdata {
		get := findLength03(tt.nums1, tt.nums2)
		if get != tt.want {
			t.Fatalf("nums1:%v, nums2:%v, want:%v, get:%v", tt.nums1, tt.nums2, tt.want, get)
		}
	}
}
