package _0_basic

func orangesRotting(grid [][]int) int {
	var dfs func(x, y, num int)

	isValid := func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}
		return false
	}

	var res int
	dfs = func(x, y, num int) {
		if !isValid(x, y) {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		if grid[x][y] == 1 { // 新鲜橘子
			grid[x][y] = num // 腐败伦茨
			res++

			dfs(x+1, y, num+1)
			dfs(x-1, y, num+1)
			dfs(x, y+1, num+1)
			dfs(x, y-1, num+1)
		}

		dfs(x+1, y, num)
		dfs(x-1, y, num)
		dfs(x, y+1, num)
		dfs(x, y-1, num)
	}

	num := 3
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				dfs(i, j, num)
			}
		}
	}
	return res
}
