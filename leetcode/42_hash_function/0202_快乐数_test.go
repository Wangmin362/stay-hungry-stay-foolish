package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/happy-number/description/

func isHappy(n int) bool {
	mm := make(map[int]struct{})

	for n > 0 {
		if _, ok := mm[n]; ok {
			return false
		}

		tmp := n
		newN := 0
		for tmp > 0 {
			mod := tmp % 10
			newN += mod * mod
			tmp /= 10
		}
		if newN == 1 {
			return true
		} else {
			mm[n] = struct{}{}
		}

		n = newN
	}

	return false
}

func TestIsHappy(t *testing.T) {
	var testdata = []struct {
		num    int
		expect bool
	}{
		{num: 19, expect: true},
		{num: 100, expect: true},
		{num: 2, expect: false},
	}

	for _, test := range testdata {
		get := isHappy(test.num)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("num:%v, expect:%v, get:%v", test.num, test.expect, get)
		}
	}

}
