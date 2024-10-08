package _1_array

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.cn/problems/n-queens/description/

func solveNQueens(n int) [][]string {
	var backtracking func(start int)
	var path [][2]int
	isValid := func(x, y int) bool {
		for _, p := range path {
			px, py := p[0], p[1]
			if px == x || py == y {
				return false
			}
			for px < n && py < n { // 检测右下角
				if px == x && py == y {
					return false
				}
				px++
				py++
			}
			px, py = p[0], p[1]
			for px < n && py >= 0 { // 检测左下角
				if px == x && py == y {
					return false
				}
				px++
				py--
			}
		}

		return true
	}

	var res [][]string
	backtracking = func(start int) {
		if len(path) == n {
			var arr []string
			for _, p := range path {
				str := strings.Builder{}
				for j := 0; j < n; j++ {
					if j == p[1] {
						str.WriteString("Q")
					} else {
						str.WriteString(".")
					}
				}
				arr = append(arr, str.String())
			}
			res = append(res, arr)
			return
		}

		for i := start; i < n; i++ {
			for j := 0; j < n; j++ {
				if !isValid(i, j) {
					continue
				}
				path = append(path, [2]int{i, j})
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtracking(0)
	return res
}

func TestSolveNQueens(t *testing.T) {
	res := solveNQueens(4)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
