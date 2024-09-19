package _1_array

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150

// 解法一：排序
/*

解法二：排序优化

如果 最小值 0 都大于等于5，那么h指数一定是最大值5
如果 倒数第二小值 1 都大于等于4，那么h指数一定是4
如果 倒数第三小值 3 都大于等于3，那么h指数一定是3
如果 倒数第四小值 5 都大于等于2，那么h指数一定是2
如果 倒数第五小值 6 都大于等于1，那么h指数一定是1
都不满足时，h指数一定是最小值0


*/

func hIndex01(citations []int) int {
	sort.Ints(citations)

	for i := len(citations); i >= 1; i-- {
		hIndex := i
		cnt := 0
		for j := 0; j < len(citations); j++ {
			if citations[j] >= hIndex {
				cnt++
			}
		}
		if cnt >= hIndex {
			return hIndex
		}
	}

	return 0
}

func hIndex02(citations []int) int {
	sort.Ints(citations)

	hi := len(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= hi {
			return hi
		}
		hi--
	}
	return hi
}

func TestHIndex(t *testing.T) {
	var testdata = []struct {
		citations []int
		want      int
	}{
		{citations: []int{3, 0, 6, 1, 5}, want: 3},
		{citations: []int{100}, want: 1},
	}

	for _, tt := range testdata {
		get := hIndex02(tt.citations)
		if get != tt.want {
			t.Fatalf("citations:%v, want:%v, get:%v", tt.citations, tt.want, get)
		}
	}
}
