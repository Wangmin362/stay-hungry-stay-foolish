package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/contains-duplicate-ii/description/?envType=study-plan-v2&envId=top-interview-150

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int][]int)
	for i, num := range nums {
		v, ok := cache[num]
		if ok {
			v = append(v, i)
			cache[num] = v
		} else {
			cache[num] = []int{i}
		}
	}

	for _, arr := range cache {
		if len(arr) <= 1 {
			continue
		}

		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] <= k {
				return true
			}
		}
	}

	return false
}

func TestContainsNearbyDuplicate(t *testing.T) {
	var testdata = []struct {
		nums []int
		k    int
		want bool
	}{
		{nums: []int{1, 2, 3, 1}, k: 3, want: true},
	}

	for _, test := range testdata {
		get := containsNearbyDuplicate(test.nums, test.k)
		if get != test.want {
			t.Fatalf("num:%v, k:%v, want:%v, get:%v", test.nums, test.k, test.want, get)
		}
	}

}
