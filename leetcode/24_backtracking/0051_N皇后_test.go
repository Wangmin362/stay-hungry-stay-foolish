package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/n-queens/description/

func solveNQueens(n int) [][]string {
	var backtracking func(n, row int)

	conflict := func(old []int, x, y int) bool {
		for oldX, oldY := range old {
			if oldX == x || oldY == y {
				return true
			} else {
				tmpX, tmpY := x, y
				for tmpX >= 0 && tmpY >= 0 { // 左上角
					if tmpX == oldX && tmpY == oldY {
						return true
					}
					tmpX--
					tmpY--
				}

				tmpX, tmpY = x, y
				for tmpX >= 0 && tmpY < n { // 右上角
					if tmpX == oldX && tmpY == oldY {
						return true
					}
					tmpX--
					tmpY++
				}
			}
		}
		return false
	}

	var res [][]string
	var path []int
	cache := make(map[int]bool)
	backtracking = func(n, row int) {
		if len(path) == n {
			tmp := make([]string, 0, n)
			for i := 0; i < n; i++ {
				str := ""
				for j := 0; j < n; j++ {
					if path[i] == j {
						str += "Q"
					} else {
						str += "."
					}
				}
				tmp = append(tmp, str)
			}

			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			exist, ok := cache[i]
			if ok && exist {
				continue
			}

			if conflict(path, row, i) {
				continue
			}

			cache[i] = true
			path = append(path, i)
			backtracking(n, row+1)
			path = path[:len(path)-1]
			cache[i] = false
		}
	}

	backtracking(n, 0)
	return res
}

func TestSolveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}
