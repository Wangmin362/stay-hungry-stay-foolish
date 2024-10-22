package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// 状态定义：dp[i][j]表示nums[i:j]的乘积
// 递推公式：dp[i][j] = dp[i+1][j-1] * nums[i] * nums[j]
func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	res := nums[0]
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if i == j {
				dp[i][j] = nums[i]
			} else if j-i == 1 {
				dp[i][j] = nums[i] * nums[j]
			} else {
				dp[i][j] = nums[i] * dp[i+1][j-1] * nums[j]
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}

// 解题思路：https://leetcode.cn/problems/maximum-product-subarray/solutions/7561/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
// 标签：动态规划
// 遍历数组时计算当前最大值，不断更新
// 令imax为当前最大值，则当前最大值为 imax = max(imax * nums[i], nums[i])
// 由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值imin，imin = min(imin * nums[i], nums[i])
// 当负数出现时则imax与imin进行交换再进行下一步计算
// 时间复杂度：O(n)
func maxProduct02(nums []int) int {
	res, imax, imin := math.MinInt, 1, 1
	for _, num := range nums {
		if num < 0 {
			imax, imin = imin, imax
		}
		imax = max(num, imax*num)
		imin = min(num, imin*num)

		res = max(res, imax)
	}
	return res
}

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct02([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct02([]int{-2, 0, -1}))
}
