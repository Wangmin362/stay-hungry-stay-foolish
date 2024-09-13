package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/house-robber-ii/description/

func robIIBacktracking(nums []int) int {
	var backtracking func(start, sum int)

	var res int
	var path []int
	backtracking = func(start, sum int) {
		res = max(res, sum)

		for i := start; i < len(nums); i++ {
			if len(path) > 0 && path[len(path)-1]+1 == i {
				continue
			}
			if len(path) > 0 && path[0] == 0 && i == len(nums)-1 {
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

// 题目分析：假设一共由四个房间，编号为0,1,2,3，我们可以分为两种情况，分别是0，1，2以及1,2,3，规则还是一样，最后取最大值即可
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	rob := func(nums []int) int {
		dp := make([]int, len(nums))
		dp[0], dp[1] = nums[0], max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}

		return dp[len(nums)-1]
	}

	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

func TestRobII(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{2, 3, 2}, want: 3},
		//{nums: []int{1, 2, 3, 1}, want: 4},
	}
	for _, tt := range testdata {
		get := robII(tt.nums)
		if get != tt.want {
			t.Fatalf("nums:%v, want:%v, get:%v", tt.nums, tt.want, get)
		}
	}
}
