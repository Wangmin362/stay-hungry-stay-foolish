package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/palindrome-partitioning/description/

func partition(s string) [][]string {

	var backtracking func(s string, startIdx, endIdx int)
	isValidStr := func(s string) bool {
		startIdx, endIdx := 0, len(s)-1
		for startIdx < endIdx {
			if s[startIdx] != s[endIdx] {
				return false
			}
			startIdx++
			endIdx--
		}
		return true
	}

	var res [][]string
	var path []string
	backtracking = func(s string, startIdx, endIdx int) {
		if startIdx > endIdx {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for idx := startIdx; idx <= endIdx; idx++ {
			str := s[startIdx : idx+1] // 当前截取的字符串
			if !isValidStr(str) {      // 当前截取的字符串是否是有效字符串，不是直接退出，完成剪枝
				continue
			}
			path = append(path, str)
			backtracking(s, idx+1, endIdx)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0, len(s)-1)
	return res
}

// 回文判断优化
func partition02(s string) [][]string {
	validStrCache := map[[2]int]bool{}
	isValidStr := func(s string, startIdx, endIdx int) bool {
		for startIdx < endIdx {
			if valid, ok := validStrCache[[2]int{startIdx, endIdx}]; ok {
				return valid
			}
			if s[startIdx] != s[endIdx] {
				validStrCache[[2]int{startIdx, endIdx}] = false
				return false
			}
			startIdx++
			endIdx--
		}
		validStrCache[[2]int{startIdx, endIdx}] = true
		return true
	}

	var res [][]string
	var path []string
	var backtracking func(s string, startIdx, endIdx int)
	backtracking = func(s string, startIdx, endIdx int) {
		if startIdx > endIdx {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for idx := startIdx; idx <= endIdx; idx++ {
			str := s[startIdx : idx+1]         // 当前截取的字符串
			if !isValidStr(s, startIdx, idx) { // 当前截取的字符串是否是有效字符串，不是直接退出，完成剪枝
				continue
			}
			path = append(path, str)
			backtracking(s, idx+1, endIdx)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0, len(s)-1)
	return res
}

func TestPartition(t *testing.T) {
	fmt.Println(partition02("aab"))
}
