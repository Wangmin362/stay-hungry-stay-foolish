package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/longest-increasing-subsequence/description/

func lengthOfLISBacktracking(nums []int) int {
	var backtracking func(start int)

	var res int
	var path []int
	backtracking = func(start int) {
		res = max(res, len(path))

		for i := start; i < len(nums); i++ {
			if len(path) > 0 && nums[i] <= path[len(path)-1] { // 必须大于最后一个数字才有意义
				continue
			}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return res
}

// 题目分析：要求求出数组的最长递增子序列，也就是说子序列的每个元素，后面一个元素都比前面一个元素大。 如果要判断一个子序列是否是递增的，那么肯定需要
// 比较数组的两个值
// 明确定义：dp[i]表示数组下标[0,i]的最长递增组序列，显然dp[i]的计算必须是从小到大，这样才能只用之前的结果
// 状态转移：dp[i]的最长递增子序列依赖于下标为[0,i-1]的每个数和nums[i]作比较，也即是说dp[i] = max(dp[j]+1, dp[i]),
// 显然只有当nums[i] > nums[j]的时候才成立这个递推公式
// 初始化：根据定义，dp[i]表示下标为[0,i]的数组中的最长递增子序列，显然做小的最长递增子序列的长度为1，也就是说初始化dp[i] = 1
// 遍历顺序：i从小到大，j从0到i-1，因为计算的dp[i]，因此只需要保证nums[i] > nums[j]即可
func lengthOfLIS0912(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ { // i从1开始计算就行，因为dp[0]肯定等于1，因为就一个元素
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// 最长的递增子序列很有可能不是考虑数组的最后一个元素，而是中间某些元素
		res = max(res, dp[i])
	}

	return res
}

func TestLengthOfLIS(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{nums: []int{10, 1, 9, 2, 5, 3, 7, 101, 18}, want: 5},
		{nums: []int{10, 1, 9, 2, 5, 3, 7, 6, 18}, want: 5},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
		{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, want: 6},
		{nums: []int{0}, want: 1},
	}

	for _, tt := range testdata {
		get := lengthOfLIS0912(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
