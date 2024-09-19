package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-interview-150

func setZeroes(matrix [][]int) {
	row, col := make(map[int]struct{}), make(map[int]struct{}) // 用于记录需要置零的行和列
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}

	// 设置行
	for i := range row {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}

	// 设置列
	for j := range col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {0, 6, 7, 8}, {9, 0, 11, 12}}

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	fmt.Println()
	setZeroes(matrix)

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
