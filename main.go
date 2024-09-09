package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var backtracking func()

	var res [][]int
	var path []int
	used := make(map[int]bool)
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	sort.Ints(nums)
	backtracking()
	return res
}

func main() {
	res := permuteUnique([]int{1, 1, 2})
	for _, r := range res {
		fmt.Println(r)
	}
}
