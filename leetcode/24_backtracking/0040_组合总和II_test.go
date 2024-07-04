package _1_array

import (
	"fmt"
	"testing"
)

// // https://leetcode.cn/problems/combination-sum/description/

func combinationSum2(candidates []int, target int) [][]int {
	var backtracking func(candidates []int, target, startIdx, sum int)

	var res [][]int
	var path []int
	backtracking = func(candidates []int, target, startIdx, sum int) {
		if target == sum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := startIdx; i < len(candidates); i++ {
			if sum > target {
				break
			}
			path = append(path, candidates[i])
			backtracking(candidates, target, i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtracking(candidates, target, 0, 0)
	return res
}

func TestSumII(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
