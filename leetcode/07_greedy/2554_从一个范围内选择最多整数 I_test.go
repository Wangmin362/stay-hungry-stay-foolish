package _0_basic

import (
	"testing"
)

// 贪心思想：记录所有选过的数字，然后从小到达选择数字
func maxCount(banned []int, n int, maxSum int) int {
	ban := make(map[int]struct{})
	for _, num := range banned {
		ban[num] = struct{}{}
	}

	res, sum := 0, 0
	for i := 1; i <= n; i++ {
		if _, ok := ban[i]; ok { // 已经使用过的数字不能再次使用
			continue
		}
		sum += i
		if sum > maxSum {
			break
		}
		res++
	}

	return res
}
func TestMaxCount(t *testing.T) {
	var testsdata = []struct {
		banned []int
		n      int
		maxSum int
		want   int
	}{
		{banned: []int{1, 6, 5}, n: 5, maxSum: 6, want: 2},
	}
	for _, tt := range testsdata {
		get := maxCount(tt.banned, tt.n, tt.maxSum)
		if get != tt.want {
			t.Fatalf("banned:%v, n:%v, maxSum:%v, want:%v, get:%v", tt.banned, tt.n, tt.maxSum, tt.want, get)
		}
	}
}
