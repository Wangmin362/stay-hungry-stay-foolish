package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/n-queens/description/

func solveNQueens(n int) [][]string {
	var backtracking func(n, deep int)

	isConflict := func(path []int, n, curr, deep int) bool { // 用于判断当前皇后是否和前面的皇后冲突
		if len(path) == 0 {
			return false
		}
		for i, j := range path {
			if curr == j { // 一定冲突
				return false
			}
			// 判断 [i, j]是否和[deep, curr]是否冲突

			// 1、向右上角走，看看是否重合，重合的话一定冲突
			x, y := deep, curr
			for x >= 0 && y < n {
				if x == i && y == j {
					return false
				}
				x--
				y++
			}

			// 2、向右下角走，看看是否重合，重合的话一定冲突
			x, y = deep, curr
			for x < n && y >= 0 {
				if x == i && y == j {
					return false
				}
				x++
				y--
			}

			// 2、向左上角走，看看是否重合，重合的话一定冲突
			x, y = deep, curr
			for x >= 0 && y >= 0 {
				if x == i && y == j {
					return false
				}
				x--
				y--
			}

			// 2、向右下角走，看看是否重合，重合的话一定冲突
			x, y = deep, curr
			for x < n && y < n {
				if x == i && y == j {
					return false
				}
				x++
				y++
			}
		}
		return true
	}

	var res [][]string
	var path []int
	backtracking = func(n, deep int) {
		if len(path) == n {
			var r []string
			for _, nn := range path {
				ss := ""
				for i := 0; i < n; i++ {
					if nn == i {
						ss += "Q"
					} else {
						ss += "."
					}
				}
				r = append(r, ss)
			}
			res = append(res, r)
			return
		}

		for i := 0; i < n; i++ {
			if isConflict(path, n, i, deep) { // 判断当前
				continue
			}
			path = append(path, i)
			backtracking(n, deep+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, 0)
	return res
}

func TestSolveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}
