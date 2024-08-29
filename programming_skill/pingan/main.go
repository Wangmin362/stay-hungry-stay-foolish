package main

import (
	"fmt"
)

const (
	n = 5 // 网格的大小为5x5
)

func findAllPath(blackX, blackY int) [][]string {
	if blackX < 0 || blackX >= n || blackY < 0 || blackY >= n {
		return nil
	}

	// 定义方向常量
	var directions = [][2]int{
		{0, 1},  // 向右
		{0, -1}, // 向左
		{1, 0},  // 向下
		{-1, 0}, // 向上
	}

	// 判断是否在网格内
	isValid := func(x, y int, visited [n][n]bool) bool {
		return x >= 0 && x < n && y >= 0 && y < n && !visited[x][y]
	}

	var backtrack func(x, y, count int, path []string, visited [n][n]bool, result *[][]string)
	// 回溯算法来寻找所有路径
	backtrack = func(x, y, count int, path []string, visited [n][n]bool, result *[][]string) {
		// 收集结果，只有当结果集覆盖了所有的格子，即所有白色格子都走了一遍，那么说明是正确的结果，其它的结果全部忽略
		if count == n*n-1 {
			*result = append(*result, append([]string(nil), path...))
			return
		}

		// 尝试所有方向
		for _, direction := range directions {
			newX, newY := x+direction[0], y+direction[1]
			if isValid(newX, newY, visited) {
				visited[newX][newY] = true
				path = append(path, fmt.Sprintf("(%d, %d)", newX, newY))
				backtrack(newX, newY, count+1, path, visited, result)
				// 回溯
				path = path[:len(path)-1]
				visited[newX][newY] = false
			}
		}
	}

	// 初始化网格和标记黑色格子的位置
	var visited [n][n]bool
	reset := func() {
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				visited[x][y] = false
			}
		}

		// 在这里设置黑色格子的位置 (例如 (2, 2) 为黑色格子)
		visited[blackX][blackY] = true // 黑色格子标记为true,表明这个格子已经走过，间接设置为不能走
	}

	var result [][]string
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if x == blackX && y == blackY { // 黑色网格点不能设置为起点
				continue
			}

			reset() // 每次重置以下网格状态

			// 设置起始位置 (例如从 (x, y) 开始)
			startX, startY := x, y
			visited[startX][startY] = true
			path := []string{fmt.Sprintf("(%d, %d)", startX, startY)}

			// 回溯查找
			backtrack(startX, startY, 1, path, visited, &result)
		}
	}

	return result
}

func main() {
	result := findAllPath(4, 1)

	// 打印所有路径
	fmt.Printf("Total number of paths: %d\n", len(result))
	for i, path := range result {
		fmt.Printf("Path %d: %v\n", i+1, path)
	}
}
