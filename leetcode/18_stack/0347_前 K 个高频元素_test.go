package _1_array

import (
	"reflect"
	"sort"
	"testing"
)

// 题目：https://leetcode.cn/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	cntMap := map[int]int{}
	for _, num := range nums {
		cntMap[num]++
	}

	var res []int
	for num := range cntMap {
		res = append(res, num)
	}
	sort.Slice(res, func(i, j int) bool {
		return cntMap[res[i]] > cntMap[res[j]]
	})

	return res[:k]
}

func TestTopKFrequents(t *testing.T) {
	var teatdata = []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			expect: []int{1, 2},
		},
	}

	for _, test := range teatdata {
		get := topKFrequent(test.nums, test.k)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("nums:%v, k:%d expect:%v, get:%v", test.nums, test.k, test.expect, get)
		}
	}
}
