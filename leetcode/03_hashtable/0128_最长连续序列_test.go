package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150

// 解法一：使用Map缓存，然后看看每个数字是否有下一个数组存在在数组当中
func longestConsecutive(nums []int) int {
	c := make(map[int]struct{})

	for _, num := range nums {
		c[num] = struct{}{}
	}

	var res int
	for _, num := range nums {
		// 有没有比当前数字更小的数字，如果有说明当前数字不是连续数字的起点，直接忽略，因为当前数字作为起点一定不是最长的
		// 否则，就把当前数字作为连续数字的起点，向后统计最大长度
		if _, ok := c[num-1]; ok {
			continue
		}

		// 以当前数字为起点找到连续的长度
		var cnt int
		_, ok := c[num]
		for ok {
			cnt++
			num++
			_, ok = c[num]
		}

		res = max(res, cnt)
	}

	return res
}

func TestLongestConsecutive(t *testing.T) {
	var testdata = []struct {
		nums []int
		want int
	}{
		{nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
	}

	for _, test := range testdata {
		get := longestConsecutive(test.nums)
		if get != test.want {
			t.Fatalf("nums:%v, want:%v, get:%v", test.nums, test.want, get)
		}
	}

}
