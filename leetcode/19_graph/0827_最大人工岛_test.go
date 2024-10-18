package _0_basic

import (
	"fmt"
	"testing"
)

func largestIsland(grid [][]int) int {
	inArea := func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}
		return false
	}

	var area func(x, y int, number int) int // 计算当前坐标所在岛屿的面积,并且标记岛屿的编号
	area = func(x, y int, number int) int {
		if !inArea(x, y) {
			return 0
		}
		if grid[x][y] != 1 { // 要么是水，要么已经遍历过
			return 0
		}
		grid[x][y] = number // 标记岛屿的编号

		return 1 + area(x-1, y, number) + area(x+1, y, number) + area(x, y-1, number) + area(x, y+1, number)
	}

	number := 2                  // 岛屿的编号从2开始
	areaMap := make(map[int]int) // 用于记录每个编号的岛屿的面积
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				ar := area(x, y, number)
				areaMap[number] = ar
				number++
			}
		}
	}

	setWaterToLand := func(x, y int) int {
		if !inArea(x, y) {
			return 0
		}
		if grid[x][y] != 0 { // 如果当前格子就是陆地，直接返回当前陆地的面积
			return areaMap[grid[x][y]]
		}

		arSet := make(map[int]struct{})
		if inArea(x-1, y) && grid[x-1][y] > 0 { //上面的格子是陆地
			arSet[grid[x-1][y]] = struct{}{}
		}
		if inArea(x+1, y) && grid[x+1][y] > 0 { //下面的格子是陆地
			arSet[grid[x+1][y]] = struct{}{}
		}
		if inArea(x, y-1) && grid[x][y-1] > 0 { //左面的格子是陆地
			arSet[grid[x][y-1]] = struct{}{}
		}
		if inArea(x, y+1) && grid[x][y+1] > 0 { //右面的格子是陆地
			arSet[grid[x][y+1]] = struct{}{}
		}

		var res int
		for nu := range arSet {
			res += areaMap[nu]
		}
		return res + 1
	}

	// 开始填海
	var res int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			ar := setWaterToLand(x, y)
			res = max(ar, res)
		}
	}

	return res
}

func TestLargestIsland(t *testing.T) {
	fmt.Println(largestIsland([][]int{{1, 1}, {1, 1}}))
}
