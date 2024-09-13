package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/pascals-triangle/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=EASY

// 对其之后可以发现规律
// [1]
// [1,1]
// [1,2,1]
// [1,3,3,1]
// [1,4,6,4,1]

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	res := make([][]int, 0, numRows)
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	for i := 3; i <= numRows; i++ {
		curr := make([]int, 0, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				curr = append(curr, 1)
				continue
			}
			n1 := res[i-2][j]
			n2 := res[i-2][j-1]
			curr = append(curr, n1+n2)
		}
		res = append(res, curr)
	}
	return res
}

func generate02(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func TestGenerate(t *testing.T) {
	res := generate02(10)
	for _, a := range res {
		fmt.Println(a)
	}

}
