package _1_array

import (
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/find-the-k-beauty-of-a-number/

func divisorSubstrings(num int, k int) int {
	nums := strconv.Itoa(num)
	if len(nums) < k {
		return 0
	}

	left, right := 0, k-1
	var res int
	for right < len(nums) {
		n, _ := strconv.Atoi(nums[left : right+1])
		if n != 0 && num/n > 0 && num%n == 0 {
			res++
		}
		left++
		right++
	}

	return res
}
func TestDivisorSubstrings(t *testing.T) {
	testdata := []struct {
		num    int
		t      int
		expect int
	}{
		{num: 240, t: 2, expect: 2},
		{num: 430043, t: 2, expect: 2},
	}

	for _, test := range testdata {
		get := divisorSubstrings(test.num, test.t)
		if get != test.expect {
			t.Errorf("num:%v, t:%v  expect:%v, get:%v", test.num, test.t, test.expect, get)
		}
	}
}
