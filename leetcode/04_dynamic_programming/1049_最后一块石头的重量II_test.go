package _0_basic

import (
	"fmt"
	"math"
	"testing"
)

// https://leetcode.cn/problems/last-stone-weight-ii/description/

// 此题可以抽象为01背包问题，题目的核心其实是要把石头分成两堆，这两对石头的重要要尽可能的相等，要整两队石头相撞之后，剩下的就是最小可能的重量。
// 因此可以把题目转换为求容量为sum/2的背包最多可以装的石头的重量（价值），此题当中，石头的重量为石头的价值。左边石头堆的重量为dp[sum/2]，
// 右边石头堆的重量为sum - dp[sum/2]，因此最小可能的重量为sum - dp[sum/2] - dp[sum/2] = sum - 2*dp[sum/2]
func lastStoneWeightII(stones []int) int {
	// dp[j]定义为容量为j的背包最多可以装的石头的价值，其实价值就是重量
	// dp[j] = max(dp[j], dp[j-weight[i]] + value[i]), 在本题当中，dp[j] = max(dp[j], dp[j - stones[i]] + stones[i])
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}

	sum := 0
	for idx := range stones {
		sum += stones[idx]
	}

	leftWeight := sum >> 1

	dp := make([]int, leftWeight+1)
	// 初始化第一行，即如果只有stones[0]的情况下，这些不同容量的背包最多可以装的石头的最大重量，只要背包的容量超过石头的重量，最多其实就是可以装
	// stones[0]的重量，因此此时就只有一个石头
	for j := stones[0]; j <= leftWeight; j++ {
		dp[j] = stones[0]
	}
	fmt.Println(dp)
	for i := 1; i < len(stones); i++ {
		for j := leftWeight; j > 0; j-- {
			if j < stones[i] { // 当前背包的容量本来就比石头的重量小，那么肯定放不进去，因此只能丢弃
				dp[j] = dp[j]
			} else {
				noput := dp[j]
				put := dp[j-stones[i]] + stones[i]
				dp[j] = int(math.Max(float64(noput), float64(put)))
			}
		}
		//fmt.Println(dp)
	}

	left := dp[leftWeight]
	right := sum - left
	return right - left
}

func lastStoneWeightII02(stones []int) int {
	// dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
	sum := 0
	for idx := range stones {
		sum += stones[idx]
	}
	mid := sum >> 1
	dp := make([]int, mid+1)
	for i := 0; i < len(stones); i++ {
		for j := mid; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
		fmt.Println(dp)
	}

	left := dp[mid]
	right := sum - left
	return right - left
}

func TestLastStoneWeightII(t *testing.T) {
	fmt.Println(lastStoneWeightII02([]int{2, 7, 4, 1, 8, 1}))
}
