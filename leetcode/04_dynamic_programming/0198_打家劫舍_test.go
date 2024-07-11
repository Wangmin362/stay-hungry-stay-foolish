package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/house-robber/description/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[j]定义为1..j个房间，可以偷到的最大金额
	// dp[j] = max(dp[j-1], dp[j-2] + num[j])
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for j := 2; j < len(nums); j++ {
		dp[j] = max(dp[j-1], dp[j-2]+nums[j])
		fmt.Println(dp)
	}

	return dp[len(nums)-1]
}

/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
/*asd*/
func rob02(nums []int) int {
	// dp[n]为n个房间中可以偷取的最大价值
	// dp[n] = max(dp[n-1], dp[n-2]+nums[n])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func TestRob(t *testing.T) {
	fmt.Println(rob02([]int{1, 2, 3, 1}))
}
