package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/n-queens-ii/description/?envType=study-plan-v2&envId=top-interview-150

func totalNQueens(n int) int {
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

	var res int
	backtracking = func(start int) {
		if len(path) == n {
			res++
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

func TestTotalNQueens(t *testing.T) {
	res := totalNQueens(4)
	fmt.Println(res)
}
