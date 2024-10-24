package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/house-robber/description/

func robBacktracking(nums []int) int {
	var backtracking func(start, sum int)

	var res int
	var path []int
	backtracking = func(start, sum int) {
		res = max(res, sum)

		for i := start; i < len(nums); i++ {
			if len(path) > 0 && path[len(path)-1]+1 == i {
				continue
			}

			path = append(path, i)
			backtracking(i+1, sum+nums[i])
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)
	return res
}

// 题目分析：当前房间头还是不偷，取决于前一个房间，似乎由一个地推关系
// 明确定义：dp[j]表示前j个房间，也就是[0,j]一共j+1个房价的最大金额，那么dp[j]的最大价值可以由dp[j-1],dp[j-2]推过来，取决于当前房间偷还是不偷
// 递推公式：dp[j] = max(dp[j-1], dp[j-2]+nums[j]), 也就是说，dp[j]的最大价值取决于前两个房价的最大价值，以及Nums[j]的价值
// 初始化：根据公司dp[0],dp[1]需要初始化。 dp[0]=nums[0], dp[1]=max(nums[0], nums[1])，当有两个房间的时候，我肯定偷价值最大的那个
// 遍历顺序：从前往后
// dp大小：[0, len(nums)-1]个房间，也就是len(nums)
func rob0912(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// 递归思路：考虑最后一个房子选还是不选  dfs(i)表示从前i个房子中获得的最大价值
// 那么：dfs(i) = max(dfs(i-1), dfs(i-2)+nums[i]) // 不选最后一个房子，那么应该考虑前i-1个房子，选择了最后一个房子应该考虑前i-2个房子
func rob(nums []int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		return max(dfs(i-1), dfs(i-2)+nums[i])
	}
	return dfs(len(nums) - 1)
}

// 记忆化搜索
func robMemory(nums []int) int {
	var dfs func(i int) int
	cache := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = -1
	}
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		cache[i] = res
		return res
	}
	return dfs(len(nums) - 1)
}

// 改为地推 f[i] = max(f[i-1], f[i-2]+nums[i])
// 给i同时加2 ==> f[i+2] = max(f[i+1], f[i]+nums[i])
func robDT(nums []int) int {
	f := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		f[i+2] = max(f[i+1], f[i]+nums[i])
	}
	return f[len(nums)+1]
}

func robDT02(nums []int) int {
	f1, f0 := 0, 0
	// f0, f1, f
	//     f0, f1, f
	for i := 0; i < len(nums); i++ {
		f := max(f1, f0+nums[i])
		f0 = f1
		f1 = f
	}
	return f1
}

func TestRob(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 1}, want: 4},
		{nums: []int{2, 7, 9, 3, 1}, want: 12},
	}
	for _, tt := range testdata {
		get := robDT02(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v get:%v", tt.nums, tt.want, get)
		}
	}
}
