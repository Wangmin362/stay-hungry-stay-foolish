package _0_basic

// https://leetcode.cn/problems/maximum-sum-circular-subarray/description/?envType=study-plan-v2&envId=top-interview-150

// 参考解题：https://leetcode.cn/problems/maximum-sum-circular-subarray/solutions/2351107/mei-you-si-lu-yi-zhang-tu-miao-dong-pyth-ilqh/?envType=study-plan-v2&envId=top-interview-150

// dp[i]为前i个子数组的最大和或者最小和
// 最大和：dp[i] = max(dp[i-1] + nums[i], nums[i])
// 最小和：dp[i] = min(dp[i-1] + nums[i], nums[i])
func maxSubarraySumCircular(nums []int) int {
	getMax := func() int {
		dp := make([]int, len(nums))
		dp[0] = nums[0]
		res := nums[0]
		for i := 1; i < len(nums); i++ {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
			res = max(res, dp[i])
		}
		return res
	}

	getMin := func() int {
		dp := make([]int, len(nums))
		dp[0] = nums[0]
		res := nums[0]
		for i := 1; i < len(nums); i++ {
			dp[i] = min(dp[i-1]+nums[i], nums[i])
			res = min(res, dp[0])
		}
		return res
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	maxSum := getMax()
	minSum := getMin()
	if minSum == sum {
		return maxSum
	}

	return max(maxSum, sum-minSum)
}
