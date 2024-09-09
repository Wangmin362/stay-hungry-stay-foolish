package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/word-search/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func exist(board [][]byte, word string) bool {
	var backtracking func(starti, startj int)

	var path []byte
	var res bool
	direction := [][2]int{
		{-1, 0}, // 向上
		{1, 0},  // 向下
		{0, -1}, // 向左
		{0, 1},  // 向右
	}
	cache := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		cache[i] = make([]bool, len(board[0]))
	}

	backtracking = func(starti, startj int) {
		if res == true {
			return
		}
		if len(path) == len(word) {
			if string(path) == word {
				res = true
			}
			return
		}

		if starti < 0 || starti >= len(board) || startj < 0 || startj >= len(board[0]) || cache[starti][startj] {
			return
		}

		for _, dir := range direction {
			newi, newj := starti+dir[0], startj+dir[1]

			cache[starti][startj] = true
			path = append(path, board[starti][startj])
			backtracking(newi, newj)
			path = path[:len(path)-1]
			cache[starti][startj] = false
		}
	}

	reset := func() {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				cache[i][j] = false
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			reset()
			backtracking(i, j)
		}
	}

	return res
}

func TestExist(t *testing.T) {
	//board := [][]byte{
	//	{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	//}

	//board := [][]byte{
	//	{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	//}
	board := [][]byte{
		{'A'},
	}
	fmt.Println(exist(board, "A"))
}
