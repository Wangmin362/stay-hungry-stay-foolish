package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/partition-equal-subset-sum/description/

// 两个和尽可能的相等，那么其中一半的集合的综合就是sum/2, 因此，只需要判断集合中是否能够装满容量为sum/2的背包，同时判断转满sum/2之后，其价值是否
// 也是等于sum/2
// 这道题的核心就是把当前提抽象为一个零一背包问题，数组中的数字其实就是不同的物品，并且物品的中间和价值其实就是相同的
// 二维数组
func canPartition01(nums []int) bool {
	if len(nums) <= 1 { // 一个以及一下的集合，肯定不能分成两个集合
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 { // 如果综合不能平分，那么一定没有这样的集合
		return false
	}

	target := sum >> 1
	// dp[i][j]定义为容量为j的最大价值dp[i][j]，因此dp[i][j] = max(dp[i-1][j], dp[i-1][j - weight[i]]) + value[i]
	// 在本题中其实就是dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i] + nums[i])
	dp := make([][]int, len(nums)) // 一共有len(nums)个物品
	dp[0] = make([]int, target+1)
	for j := 1; j <= target && nums[0] <= j; j++ {
		dp[0][j] = nums[0]
	}
	fmt.Println(dp[0])

	for i := 1; i < len(nums); i++ {
		dp[i] = make([]int, target+1)
		for j := 1; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				noput := dp[i-1][j]
				put := dp[i-1][j-nums[i]] + nums[i]
				dp[i][j] = int(math.Max(float64(noput), float64(put)))
			}
		}
		fmt.Println(dp[i])
	}

	return dp[len(nums)-1][target] == target
}

// 二位数组，从后向前遍历，先遍历物品，在遍历价值
func canPartition02(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 { // 不能整除一定不能平分
		return false
	}

	target := sum >> 1

	// dp[i][j] 定义为物品0..i放入到容量为j的背包中的最大价值，此题目中物品的重量其实就是物品的最大价值
	dp := make([][]int, len(nums)) // 一共有len(nums)个物品
	dp[0] = make([]int, target+1)
	for j := 1; j <= target && j >= nums[0]; j++ {
		dp[0][j] = nums[0]
	}
	fmt.Println(dp[0])

	for i := 1; i < len(nums); i++ { // 从第一个物品开始遍历
		dp[i] = make([]int, target+1)
		for j := target; j > 0; j-- { // 从后向前遍历，因为当前的dp[i][j]只取决于上一样的状态，所以从后往前遍历没有问题
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				noput := dp[i-1][j]                 // 不放入物品i
				put := dp[i-1][j-nums[i]] + nums[i] // 放入物品i, 只有背包的容量为j-nums[i]，才能保证一定能放进入物品i
				dp[i][j] = int(math.Max(float64(noput), float64(put)))
			}
		}
		fmt.Println(dp[i])
	}

	return dp[len(nums)-1][target] == target
}

// 二位数组  先遍历价值，在遍历物品
func canPartition03(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	target := sum >> 1

	dp := make([][]int, len(nums))
	for idx := range nums { // 先分配数组的存储空间
		dp[idx] = make([]int, target+1)
	}

	// 第一列，肯定是0， 不需要初始化
	for j := 1; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if i == 0 {
				if j >= nums[i] {
					dp[i][j] = nums[0]
				}
			} else if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				noput := dp[i-1][j]
				put := dp[i-1][j-nums[i]] + nums[i]
				dp[i][j] = int(math.Max(float64(noput), float64(put)))
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(dp[i])
	}

	return dp[len(nums)-1][target] == target
}

// 一维数组 dp[j] = map(dp[j], dp[j-weight[i]] + value[i])
// 此时遍历价值只能从后向前遍历，只有这样物品的价值才可能被加一次，否则会被加两次
func canPartition04(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}

	target := sum >> 1
	dp := make([]int, target+1)
	for j := 1; j <= target && j >= nums[0]; j++ { // 初始化第一行
		dp[j] = nums[0]
	}
	fmt.Println(dp)
	for i := 1; i < len(nums); i++ {
		for j := target; j > 0; j-- {
			if j >= nums[i] {
				noput := dp[j]                 // 不放入i物品
				put := dp[j-nums[i]] + nums[i] // 放入i物品
				dp[j] = int(math.Max(float64(noput), float64(put)))
			}
		}
		fmt.Println(dp)
	}

	return dp[target] == target
}

func TestCanPartition(t *testing.T) {
	canPartition04([]int{1, 5, 11, 5})
	canPartition04([]int{2, 2, 3, 5})
}
