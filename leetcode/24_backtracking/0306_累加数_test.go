package _1_array

import (
	"fmt"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/additive-number/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func isAdditiveNumber(num string) bool {
	var backtracking func(start int)

	var res bool
	var path []int
	backtracking = func(start int) {
		if res {
			return
		}
		if start >= len(num) {
			if len(path) >= 3 {
				fmt.Println(path)
				res = true
			}
			return
		}

		for i := start; i < len(num); i++ {
			if len(num[start:i+1]) > 1 && num[start] == '0' {
				continue
			}
			n, _ := strconv.Atoi(num[start : i+1])
			if len(path) >= 2 && path[len(path)-1]+path[len(path)-2] != n {
				continue
			}

			path = append(path, n)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return res
}

func TestIsAdditiveNumber(t *testing.T) {
	//fmt.Println(isAdditiveNumber("112358"))
	//fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("1023"))
}
