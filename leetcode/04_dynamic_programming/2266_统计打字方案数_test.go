package _0_basic

import (
	"fmt"
	"testing"
)

// 回溯算法
func countTextsBacktracking(pressedKeys string) int {
	var backtracking func(start int)

	var res int
	var path []string
	backtracking = func(start int) {
		if len(path) > 0 && start >= len(pressedKeys) {
			res++
			//fmt.Println(path)
		}

		for i := start; i < len(pressedKeys); i++ {
			isValid := true
			for j := start + 1; j <= i; j++ { // 每个字符相同才行
				if pressedKeys[j] != pressedKeys[j-1] {
					isValid = false
					break
				}
			}
			if !isValid {
				continue
			}
			path = append(path, pressedKeys[start:i+1])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}

	}

	backtracking(0)
	return res
}

func TestCountTextsBacktracking(t *testing.T) {
	fmt.Println(countTextsBacktracking("222222222222222222222222222222222222"))
}
