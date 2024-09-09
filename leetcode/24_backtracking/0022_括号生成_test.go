package _1_array

import (
	"container/list"
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/generate-parentheses/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func generateParenthesis(n int) []string {
	var backtracking func()

	isValidParn := func(parns []byte) bool {
		stack := list.New()
		for _, p := range parns {
			if p == ')' {
				if stack.Len() == 0 {
					return false
				}
				x := stack.Remove(stack.Back()).(byte)
				if x != '(' {
					return false
				}
			} else {
				stack.PushBack(p)
			}
		}
		if stack.Len() != 0 {
			return false
		}
		return true
	}

	var res []string
	var path []byte
	validPar := []byte{'(', ')'}
	backtracking = func() {
		if len(path) == 2*n {
			if isValidParn(path) {
				res = append(res, string(path))
			}
			return
		}

		for i := 0; i < len(validPar); i++ {
			path = append(path, validPar[i])
			backtracking()
			path = path[:len(path)-1]
		}
	}

	backtracking()
	return res
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
