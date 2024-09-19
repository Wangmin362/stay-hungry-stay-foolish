package _0_basic

import (
	"fmt"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for l <= r && t <= b {
		for i := l; i <= r; i++ { // 从左到右
			res = append(res, matrix[t][i])
		}
		t++
		if b < t {
			break
		}

		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if b < t {
			break
		}

		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	fmt.Println(res)
	return res
}

// 指定上下左右的边界，然后按照上，右，下左的方式遍历
func spiralOrder02(matrix [][]int) []int {
	t, r, b, l := 0, len(matrix[0])-1, len(matrix)-1, 0
	var res []int
	for l <= r && t <= b {
		for i := l; i <= r; i++ { // 遍历上
			res = append(res, matrix[t][i])
		}
		t++ // 缩小上边界
		if b < t {
			break
		}

		for i := t; i <= b; i++ { // 遍历右
			res = append(res, matrix[i][r])
		}
		r-- // 缩小有边界
		if l > r {
			break
		}

		for i := r; i >= l; i-- { // 遍历下
			res = append(res, matrix[b][i])
		}
		b-- // 缩小下边界
		if b < t {
			break
		}

		for i := b; i >= t; i-- { // 遍历左
			res = append(res, matrix[i][l])
		}
		l++ // 缩小左边界
		if l > r {
			break
		}
	}

	return res
}

func TestSpiralOrder(t *testing.T) {
	testdata := []struct {
		n      [][]int
		expect []int
	}{
		// {n: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{n: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		// {n: [][]int{{1, 2, 3, 4}, {4, 5, 6, 9}}, expect: []int{1, 2, 3, 4, 9, 6, 5, 4}},
	}

	for _, test := range testdata {
		get := spiralOrder02(test.n)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("n:%v, expect:%v, get:%v", test.n, test.expect, get)
		}
	}
}
