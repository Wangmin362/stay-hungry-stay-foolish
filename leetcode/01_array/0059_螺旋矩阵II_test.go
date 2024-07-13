package _1_array

import (
	"fmt"
	"reflect"
	"testing"
)

// 题目：https://leetcode.cn/problems/spiral-matrix-ii/description/

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1 // 左、右、上、下边界初始化
	num := 0
	target := n * n
	for num < target {
		for i := l; i <= r; i++ { // 从左到右
			num++
			res[t][i] = num
		}
		t++ // 缩小上边界

		for i := t; i <= b; i++ { // 从上到下
			num++
			res[i][r] = num
		}
		r-- // 缩小有边界

		for i := r; i >= l; i-- { // 从右到左
			num++
			res[b][i] = num
		}
		b-- // 缩小底部边界

		for i := b; i >= t; i-- { // 从下到上
			num++
			res[i][l] = num
		}
		l++
	}
	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}
	return res
}

func TestGenerateMatrix(t *testing.T) {
	testdata := []struct {
		n      int
		expect [][]int
	}{
		{n: 3, expect: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, test := range testdata {
		get := generateMatrix(test.n)
		if !reflect.DeepEqual(get, test.expect) {
			t.Errorf("n:%v, expect:%v, get:%v", test.n, test.expect, get)
		}
	}
}
