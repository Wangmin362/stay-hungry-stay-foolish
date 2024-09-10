package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/matchsticks-to-square/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func makesquare(matchsticks []int) bool {
	var backtracking func(edge int)

	var res bool
	var path [4]int

	cache := make([]bool, len(matchsticks))
	cacheValid := func() bool {
		for i := 0; i < len(cache); i++ {
			if !cache[i] {
				return false
			}
		}
		return true
	}

	backtracking = func(edge int) {
		if res {
			return
		}
		if cacheValid() {
			valid := true
			for i := 1; i < 4; i++ {
				if path[i] != path[i-1] {
					valid = false
					break
				}
			}
			if valid {
				res = true
			}
			return
		}

		for i := 0; i < len(matchsticks); i++ {
			if cache[i] { // 当前火柴已经使用过了
				continue
			}

			cache[i] = true
			path[edge%4] += matchsticks[i]
			backtracking(edge + 1)
			path[edge%4] -= matchsticks[i]
			cache[i] = false
		}
	}

	backtracking(0)
	return res
}

func TestMakesquare(t *testing.T) {
	//fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	//fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	//fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
	fmt.Println(makesquare([]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}))
}
