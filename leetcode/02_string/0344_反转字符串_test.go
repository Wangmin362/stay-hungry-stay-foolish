package _1_array

import (
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/reverse-string/description/

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func TestReverseString(t *testing.T) {
	var teatdata = []struct {
		s      []byte
		expect []byte
	}{
		//	todo
	}

	for _, test := range teatdata {
		reverseString(test.s)
		if !reflect.DeepEqual(test.s, test.expect) {
			t.Errorf("expect:%v, get:%v", test.expect, test.s)
		}
	}
}
