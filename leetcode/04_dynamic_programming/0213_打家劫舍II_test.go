package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/house-robber-ii/description/

// 把问题展开为打家劫舍问题 1, 2, 3, 1可以展开为1, 2, 3或者2, 3, 1
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	rob := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		// dp[i]定义为0..i个房子偷取现金的最大值
		dp := make([]int, len(nums))
		dp[0], dp[1] = nums[0], max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}

		return dp[len(nums)-1]
	}

	case1 := rob(nums[:len(nums)-1])
	case2 := rob(nums[1:])
	return max(case1, case2)
}

func TestRobII(t *testing.T) {
	fmt.Println(robII([]int{1}))
}
