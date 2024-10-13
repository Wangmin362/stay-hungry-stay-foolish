package _0_basic

import (
	"sort"
	"testing"
)

// 贪心思想：统计每种数字出现的频率，并且按照从小到大排序，优先移除频率最小的数字，最后剩下的就是最优解
func findLeastNumOfUniqueInts(arr []int, k int) int {
	c := make(map[int]int, 128)
	for _, num := range arr {
		c[num]++
	}

	freq := make([]int, 0, len(c)) // 统计每个数字的频率，并且按照从小到大排序
	for _, cnt := range c {
		freq = append(freq, cnt)
	}
	sort.Ints(freq)

	res := len(freq)
	for i := 0; i < len(freq) && k > 0; i++ {
		if k >= freq[i] {
			res--
			k -= freq[i]
		}
	}

	return res
}
func TestFindLeastNumOfUniqueInts(t *testing.T) {
	var testsdata = []struct {
		arr  []int
		k    int
		want int
	}{
		//{arr: []int{5, 5, 4}, k: 1, want: 1},
		{arr: []int{9, 17, 11, 19, 4, 22, 27, 15, 24, 30, 45, 11, 17, 37, 37}, k: 8, want: 4},
	}
	for _, tt := range testsdata {
		get := findLeastNumOfUniqueInts(tt.arr, tt.k)
		if get != tt.want {
			t.Fatalf("arr:%v, k:%v, want:%v, get:%v", tt.arr, tt.k, tt.want, get)
		}
	}
}
