package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/ambiguous-coordinates/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func ambiguousCoordinates(s string) []string {
	var backtracking func(start int)

	isValid := func(str string) bool {
		if len(str) > 1 && str[0] == '0' {
			return false
		}
		return true
	}

	isValidFloat := func(str string) bool {
		if len(str) > 1 && str[0] == '0' && str[len(str)-1] == '0' {
			return false
		}
		if str[len(str)-1] == '0' {
			return false
		}
		return true
	}

	var res []string
	var path []string
	backtracking = func(start int) {
		if len(path) >= 2 && start >= len(s)-1 { // 收集结果
			switch len(path) {
			case 2:
				if isValid(path[0]) && isValid(path[1]) {
					res = append(res, fmt.Sprintf("(%s, %s)", path[0], path[1]))
				}
			case 3:
				if isValidFloat(path[1]) && isValid(path[0]) && isValid(path[2]) {
					res = append(res, fmt.Sprintf("(%s.%s, %s)", path[0], path[1], path[2]))
				}
				if isValid(path[0]) && isValid(path[1]) && isValidFloat(path[2]) {
					res = append(res, fmt.Sprintf("(%s, %s.%s)", path[0], path[1], path[2]))
				}
			case 4:
				if isValidFloat(path[1]) && isValidFloat(path[3]) && isValid(path[0]) && isValid(path[2]) {
					res = append(res, fmt.Sprintf("(%s.%s, %s.%s)", path[0], path[1], path[2], path[3]))
				}
				return
			}
		}

		for i := start; i < len(s)-1; i++ {
			ss := s[start : i+1]
			if len(path) == 4 && i != len(s)-1 {
				break
			}

			path = append(path, ss)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(1)
	return res
}

func TestAmbiguousCoordinates(t *testing.T) {
	fmt.Println(ambiguousCoordinates("(0010)"))
}
