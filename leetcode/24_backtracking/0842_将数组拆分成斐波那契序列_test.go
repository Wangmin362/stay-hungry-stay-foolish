package _1_array

import (
	"fmt"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/split-array-into-fibonacci-sequence/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func splitIntoFibonacci(num string) []int {
	var backtracking func(start int)

	var res []int
	var done bool
	var path []int
	backtracking = func(start int) {
		if done {
			return
		}
		if start >= len(num) {
			if len(path) >= 3 {
				isValid := true
				for i := 2; i < len(path); i++ {
					if path[i] != path[i-2]+path[i-1] {
						isValid = false
						break
					}
				}
				if isValid {
					tmp := make([]int, len(path))
					copy(tmp, path)
					res = tmp
					done = true
				}
			}
			return
		}

		for i := start; i < len(num); i++ {
			if i > start && num[start] == '0' { // 前导零不合法
				continue
			}
			if len(num)-i+1+len(path) < 3 { // 必须满足3个以上
				break
			}

			n, _ := strconv.Atoi(num[start : i+1])
			if n > 1<<31 {
				continue
			}
			if len(path) >= 2 && n != path[len(path)-1]+path[len(path)-2] {
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

func TestSplitIntoFibonacci(t *testing.T) {
	fmt.Println(splitIntoFibonacci("74912134825162255812723932620170946950766784234934"))
}
