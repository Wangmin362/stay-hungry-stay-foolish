package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/path-with-maximum-gold/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func getMaximumGold(grid [][]int) int {
	var backtracking func(si, sj int)
	directions := [][]int{
		{-1, 0}, // 向上
		{1, 0},  // 向下
		{0, -1}, // 向左
		{0, 1},  // 向右
	}

	cache := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		cache[i] = make([]bool, len(grid[0]))
	}
	cacheReset := func() {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				cache[i][j] = false
			}
		}
	}

	res := 0
	cur := 0
	//var path []int
	backtracking = func(si, sj int) {
		res = max(res, cur)
		//fmt.Println(path)

		if si < 0 || si >= len(grid) || sj < 0 || sj >= len(grid[0]) || cache[si][sj] || grid[si][sj] == 0 {
			return
		}

		for _, dir := range directions {
			//path = append(path, grid[si][sj])
			cur += grid[si][sj]
			cache[si][sj] = true
			backtracking(si+dir[0], sj+dir[1])
			cache[si][sj] = false
			cur -= grid[si][sj]
			//path = path[:len(path)-1]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			cacheReset()
			backtracking(i, j)
		}
	}

	return res
}

func TestGetMaximumGold(t *testing.T) {
	gird := [][]int{
		{12, 23, 0, 0, 18, 18, 34}, {0, 0, 0, 0, 35, 0, 0}, {21, 20, 0, 0, 0, 18, 38}, {36, 18, 4, 30, 2, 8, 20}, {0, 0, 33, 36, 28, 0, 11}, {32, 13, 0, 15, 0, 0, 40},
	}
	fmt.Println(getMaximumGold(gird))
}
