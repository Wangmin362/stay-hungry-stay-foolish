package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/combination-sum/description/

func combinationSum(candidates []int, target int) [][]int {
	var backtracking func(candidates []int, startIdx, target, sum int) // k表示当前需要挑选出k个数组

	var res [][]int
	var path []int
	backtracking = func(candidates []int, startIdx, target, sum int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := startIdx; i < len(candidates); i++ { // 因为数字可以重复去，所以每次for循环的集合就是全部数组
			if sum > target {
				break
			}
			path = append(path, candidates[i])
			backtracking(candidates, i, target, sum+candidates[i])
			path = path[:len(path)-1]
		}

	}

	backtracking(candidates, 0, target, 0)

	return res
}

func TestSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
