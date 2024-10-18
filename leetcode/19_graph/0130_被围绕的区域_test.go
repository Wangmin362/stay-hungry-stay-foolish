package _0_basic

import (
	"fmt"
	"testing"
)

// 1, 先便利四周，然后吧所有的O都标记为B，然后把所有的O标记为X，最后还原所有的B为O
func solve(board [][]byte) {
	inArea := func(x, y int) bool {
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
			return true
		}
		return false
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !inArea(x, y) {
			return
		}
		if board[x][y] != 'O' { // 要么是X，要么是B
			return
		}
		board[x][y] = 'B'

		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	// 遍历四周，把所有的O标记为B
	for y := 0; y < len(board[0]); y++ { // 处理此一行
		if board[0][y] == 'O' {
			dfs(0, y)
		}
		if board[len(board)-1][y] == 'O' { // 处理最后一行
			dfs(len(board)-1, y)
		}
	}
	for x := 0; x < len(board); x++ {
		if board[x][0] == 'O' { // 处理第一列
			dfs(x, 0)
		}
		if board[x][len(board[0])-1] == 'O' { // 处理最后一列
			dfs(x, len(board[0])-1)
		}
	}

	// 把所有的O设置为X
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == 'O' {
				board[x][y] = 'X'
			}
		}
	}

	// 把所有的B还原为O
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == 'B' {
				board[x][y] = 'O'
			}
		}
	}
}

func TestSlove(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
