package _1_array

import (
	"fmt"
	"testing"
)

// 回溯
func minCutBacktracking(s string) int {
	var backtracking func(start int)

	m := make(map[[2]int]bool)
	isValid := func(start, end int) bool {
		if isValid, ok := m[[2]int{start, end}]; ok {
			return isValid
		}
		for start < end {
			if s[start] != s[end] {
				m[[2]int{start, end}] = false
				return false
			}
			start++
			end--
		}

		m[[2]int{start, end}] = true
		return true
	}

	minCnt := len(s) - 1 // 最长分割就是每个字符
	var path []string
	backtracking = func(start int) {
		if start >= len(s) {
			minCnt = min(minCnt, len(path)-1)
			return
		}
		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if !isValid(start, i) {
				continue
			}
			path = append(path, str)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return minCnt
}

func TestMinCut(t *testing.T) {
	fmt.Println(minCutBacktracking("acdeab"))
}
