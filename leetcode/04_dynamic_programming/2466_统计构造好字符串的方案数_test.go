package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/count-ways-to-build-good-strings/description/

// 题目分析：可以抽象为一个爬楼梯问题，需要爬到[low, high]，每次可以选择爬zero或者one步数，问爬到[low, high]一共有多少布
// 递推公式：dp[i] = dp[i-zero] + dp[i-one]
// 初始化：比较zero以及one的大小，小于最小数字的dp[i]=0
// 遍历顺序：从小到大
func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	mod := 1_000_000_007

	var res int
	for i := 1; i <= high; i++ {
		if i >= zero && i >= one {
			dp[i] = (dp[i-zero] + dp[i-one]) % mod
		} else if i >= zero {
			dp[i] = dp[i-zero]
		} else if i >= one {
			dp[i] = dp[i-one]
		}
		if i >= low {
			res = (res + dp[i]) % mod
		}
	}

	return res
}

func TestCountGoodStrings(t *testing.T) {
	var testdata = []struct {
		low  int
		high int
		zero int
		one  int
		want int
	}{
		{low: 3, high: 3, zero: 1, one: 1, want: 8},
		{low: 2, high: 3, zero: 1, one: 2, want: 5},
		{low: 1, high: 100000, zero: 1, one: 1, want: 215447031},
	}

	for _, tt := range testdata {
		get := countGoodStrings(tt.low, tt.high, tt.zero, tt.one)
		if get != tt.want {
			t.Fatalf("low:%v, high:%v, zero:%v, one:%v, want:%v, get:%v", tt.low, tt.high, tt.zero, tt.one, tt.want, get)
		}
	}

}
