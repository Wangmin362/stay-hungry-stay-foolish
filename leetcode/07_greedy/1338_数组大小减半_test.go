package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：统计每个数字出现的频率，然后删除频率最高的数字即可
func minSetSize(arr []int) int {
	c := make(map[int]int, 128)
	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}

	freq := make([]int, 0, len(c))
	for _, cnt := range c {
		freq = append(freq, cnt)
	}
	sort.Ints(freq)

	mid := len(arr) >> 1
	if len(arr) == 1 {
		mid += 1
	}

	var res int
	for i := len(freq) - 1; i >= 0; i-- {
		res++
		mid -= freq[i]
		if mid <= 0 {
			break
		}
	}

	return res
}

func TestMinSetSize(t *testing.T) {
	var testsdata = []struct {
		arr  []int
		want int
	}{
		{arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, want: 2},
		{arr: []int{1, 9}, want: 1},
	}
	for _, tt := range testsdata {
		get := minSetSize(tt.arr)
		if get != tt.want {
			t.Fatalf("arr:%v,  want:%v, get:%v", tt.arr, tt.want, get)
		}
	}
}
