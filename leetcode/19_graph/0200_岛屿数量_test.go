package _0_basic

import (
	"fmt"
	"testing"
)

func numIslands(grid [][]byte) int {
	inArea := func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}

		return false
	}

	var dfs func(x, y int)

	dfs = func(x, y int) {
		if !inArea(x, y) {
			return
		}
		if grid[x][y] != '1' { // 要么是水，要么已经遍历过了
			return
		}
		grid[x][y] = 2

		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	var res int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == '1' {
				dfs(x, y)
				res++
			}
		}
	}

	return res
}

func TestNumIslands(t *testing.T) {
	ar := [][]byte{
		{'1', '1', '1', '1', 0},
		{'1', '1', 0, '1', 0},
		{'1', '1', 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(ar))
}
