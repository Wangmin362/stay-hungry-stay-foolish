package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：先排序，然后选择每次选择幸福指数最大的孩子
func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	var res int
	for i := len(happiness) - 1; i >= 0 && k > 0; i-- {
		happy := happiness[i] - (len(happiness) - 1 - i)
		if happy <= 0 {
			break
		}
		res += happy
		k--
	}

	return int64(res)
}

func TestMaximumHappinessSum(t *testing.T) {
	var testsdata = []struct {
		happiness []int
		k         int
		want      int64
	}{
		{happiness: []int{1, 2, 3}, k: 2, want: 4},
		{happiness: []int{1, 1, 1, 1, 1, 1}, k: 5, want: 1},
	}
	for _, tt := range testsdata {
		get := maximumHappinessSum(tt.happiness, tt.k)
		if get != tt.want {
			t.Fatalf("happiness:%v, k:%v, want:%v, get:%v", tt.happiness, tt.k, tt.want, get)
		}
	}
}
