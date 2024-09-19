package _0_basic

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

func isHappy01(n int) bool {
	happy := func(n int) int {
		sum := 0
		for n >= 10 {
			mod := n % 10
			sum += mod * mod
			n /= 10
		}
		sum += n * n
		return sum
	}

	cache := map[int]struct{}{}
	for {
		n = happy(n)
		if n == 1 {
			return true
		} else {
			if _, ok := cache[n]; ok {
				return false
			} else {
				cache[n] = struct{}{}
			}
		}
	}
}

func isHappy03(n int) bool {
	cache := make(map[int]struct{})
	powSum := func(n int) int {
		var sum int
		for n > 0 {
			mod := n % 10
			n /= 10
			sum += mod * mod
		}
		return sum
	}
	for {
		res := powSum(n)
		if res == 1 {
			return true
		}
		if _, ok := cache[res]; ok {
			return false
		}
		cache[n] = struct{}{}
		n = res
	}
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
		get := isHappy03(test.num)
		if !reflect.DeepEqual(get, test.expect) {
			t.Fatalf("num:%v, expect:%v, get:%v", test.num, test.expect, get)
		}
	}

}
