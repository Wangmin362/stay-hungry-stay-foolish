package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：每次购买最便宜的雪糕，就能买到最多的数量
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var res int
	for i := 0; i < len(costs) && coins > 0; i++ {
		if coins >= costs[i] { // 剩余的钱可以买当前雪糕
			res++
			coins -= costs[i]
		}
	}

	return res
}

func TestMaxIceCream(t *testing.T) {
	var testsdata = []struct {
		costs []int
		coins int
		want  int
	}{
		{costs: []int{1, 3, 2, 4, 1}, coins: 7, want: 4},
	}
	for _, tt := range testsdata {
		get := maxIceCream(tt.costs, tt.coins)
		if get != tt.want {
			t.Fatalf("costs:%v, coins:%v, want:%v, get:%v", tt.costs, tt.coins, tt.want, get)
		}
	}
}
