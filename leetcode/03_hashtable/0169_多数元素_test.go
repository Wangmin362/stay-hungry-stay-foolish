package _0_basic

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

// 解法一：使用map记录每个元素出现的次数
func majorityElement(nums []int) int {
	mm := make(map[int]int)
	cnt := len(nums) >> 1
	for _, num := range nums {
		mm[num] += 1
		if mm[num] > cnt {
			return num
		}
	}
	return 0
}

// 摩尔投票法，相当于多数投票的人 - 其余所有人一定是大于等于1的
func majorityElement01(nums []int) int {
	winner, cnt := nums[0], 0
	for _, num := range nums {
		if num == winner {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			winner = num
			cnt++
		}
	}

	return winner
}

func TestMajorityElement(t *testing.T) {
	var testdata = []struct {
		num    []int
		expect int
	}{
		{num: []int{2, 1, 2}, expect: 2},
		{num: []int{6, 5, 5}, expect: 5},
	}

	for _, test := range testdata {
		get := majorityElement01(test.num)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("num:%v, expect:%v, get:%v", test.num, test.expect, get)
		}
	}

}
