package _1_array

import (
	"fmt"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return nil
	}
	res := make([]int, 0, len(array)*len(array[0]))
	l, r, t, b := 0, len(array[0])-1, 0, len(array)-1
	for l <= r && t <= b {
		for i := l; i <= r; i++ { // 从左到右
			res = append(res, array[t][i])
		}
		t++
		if b < t {
			break
		}

		for i := t; i <= b; i++ {
			res = append(res, array[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, array[b][i])
		}
		b--
		if b < t {
			break
		}

		for i := b; i >= t; i-- {
			res = append(res, array[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	fmt.Println(res)
	return res
}

func TestSpiralArrayOrder(t *testing.T) {
	testdata := []struct {
		n      [][]int
		expect []int
	}{
		// {n: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{n: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		// {n: [][]int{{1, 2, 3, 4}, {4, 5, 6, 9}}, expect: []int{1, 2, 3, 4, 9, 6, 5, 4}},
	}

	for _, test := range testdata {
		get := spiralArray(test.n)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("n:%v, expect:%v, get:%v", test.n, test.expect, get)
		}
	}
}
