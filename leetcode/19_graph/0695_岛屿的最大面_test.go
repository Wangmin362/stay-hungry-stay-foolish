package _0_basic

func maxAreaOfIsland(grid [][]int) int {
	inArea := func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}
		return false
	}

	var area func(x, y int) int
	area = func(x, y int) int { // 返回以x,y坐标所在岛屿的面积
		if !inArea(x, y) { // 如果超出边界了，直接返回0
			return 0
		}

		if grid[x][y] != 1 { // 如果当前坐标不是岛屿直接返回0
			return 0
		}
		grid[x][y] = 2 // 标记遍历过的岛屿

		return 1 + area(x-1, y) + area(x+1, y) + area(x, y-1) + area(x, y+1)
	}

	var res int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 { // 只有当前坐标是岛屿的情况下才遍历
				ar := area(x, y)
				res = max(ar, res)
			}
		}
	}

	return res
}
