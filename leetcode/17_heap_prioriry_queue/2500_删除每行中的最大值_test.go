package _0_basic

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/delete-greatest-value-in-each-row/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

func deleteGreatestValue(grid [][]int) int {
	for _, g := range grid {
		sort.Ints(g)
	}

	var res int
	for i := 0; i < len(grid[0]); i++ {
		m := math.MinInt32
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > m {
				m = grid[j][i]
			}
		}
		res += m
	}
	return res
}
