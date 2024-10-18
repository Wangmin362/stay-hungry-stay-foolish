package _0_basic

func islandPerimeter(grid [][]int) int {
	inArea := func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}
		return false
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if !inArea(x, y) {
			return 1
		}

		if grid[x][y] == 0 {
			return 1
		}
		if grid[x][y] != 1 { // 说明已经遍历过了
			return 0
		}
		grid[x][y] = 2
		return dfs(x-1, y) + dfs(x+1, y) + dfs(x, y-1) + dfs(x, y+1)
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				return dfs(x, y)
			}
		}
	}
	return 0
}
