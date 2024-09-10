package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/count-increasing-quadruplets/description/?envType=daily-question&envId=2024-09-10

// 回溯
func countQuadruplets(nums []int) int64 {
	if len(nums) < 4 {
		return 0
	}

	var backtracking func(start int)

	var res int
	var path []int
	backtracking = func(start int) {
		if len(path) == 4 {
			if nums[path[0]] < nums[path[2]] && nums[path[2]] < nums[path[1]] && nums[path[1]] < nums[path[3]] {
				res++
			}
			return
		}

		for i := start; i < len(nums); i++ {
			if len(nums)-i+1 < 4-len(path) {
				break
			}

			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}

	}

	backtracking(0)
	return int64(res)
}

func TestXXX(t *testing.T) {
	fmt.Println(countQuadruplets([]int{1, 2, 3, 4}))
	fmt.Println(countQuadruplets([]int{1, 3, 2, 4, 5}))
}
