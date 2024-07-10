package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/target-sum/description/

// 也可以抽象为01背包问题，每个元素只能取一次。把数组分为两个组，分别是left, right。根据题目的意思left - right = 3 并且left + right = sum
// 那么left - (sum -left) = 3，即2left - sum = 3，即left = (sum + 3) / 2，因此left = (sum + target) /2
func findTargetSumWays(nums []int, target int) int {
	// dp[j]定义为容量为j的背包，有多少种方法可以装满，
	// 根据上面的定义，dp[j] = dp[j-1]+dp[j-2]+...dp[0]
	// dp[j] += dp[j - nums[i]]
	sum := 0
	for idx := range nums {
		sum += nums[idx]
	}
	if target > sum {
		return 0
	}
	if target < 0 && -target > sum {
		return 0
	}

	if (sum+target)%2 == 1 { // 说明无法拼凑
		return 0
	}

	left := (sum + target) / 2
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
		fmt.Println(dp)
	}

	return dp[left]
}

func findTargetSumWays02(nums []int, target int) int {
	// left + right = sum
	//                         ==>  left - (sum - left) = target => left = (sum + target)/2
	// left - right = target
	sum := 0
	for idx := range nums {
		sum += nums[idx]
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	if math.Abs(float64(target)) > float64(sum) {
		return 0
	}
	left := (sum + target) >> 1
	// dp[j] += dp[j - nums[i]]
	dp := make([]int, left+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := left; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[left]
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays02([]int{1000}, -1000))
}
