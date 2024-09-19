package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/valid-sudoku/description/?envType=study-plan-v2&envId=top-interview-150

func isValidSudoku(board [][]byte) bool {
	cache := make([]bool, 9)
	reset := func() {
		for i := 0; i < 9; i++ {
			cache[i] = false
		}
	}

	// 检查行
	for i := 0; i < 9; i++ {
		reset()
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0' - 1
			if cache[num] {
				return false
			}
			cache[num] = true
		}
	}

	// 检查列
	for j := 0; j < 9; j++ {
		reset()
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0' - 1
			if cache[num] {
				return false
			}
			cache[num] = true
		}
	}

	start := [][2]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}
	axis := [][]int{
		{0, 0}, {0, 1}, {0, 2},
		{1, 0}, {1, 1}, {1, 2},
		{2, 0}, {2, 1}, {2, 2},
	}
	// 检查方格
	for _, st := range start {
		reset()
		for _, ax := range axis {
			i, j := st[0]+ax[0], st[1]+ax[1]
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0' - 1
			if cache[num] {
				return false
			}
			cache[num] = true
		}
	}

	return true
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}
