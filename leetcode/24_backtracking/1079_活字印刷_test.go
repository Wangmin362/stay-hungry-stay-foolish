package _1_array

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/letter-tile-possibilities/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM

func numTilePossibilities(tiles string) int {
	var backtracking func()

	ts := make([]int, len(tiles))
	for idx, s := range tiles {
		ts[idx] = int(s)
	}

	var res int
	var path []int
	used := make([]bool, len(tiles))
	backtracking = func() {
		if len(path) > 0 {
			res++
		}

		for i := 0; i < len(tiles); i++ {
			if used[i] {
				continue
			}
			if i > 0 && ts[i] == ts[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, ts[i])
			backtracking()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	sort.Ints(ts)
	backtracking()
	return res
}

func TestNumTilePossibilities(t *testing.T) {
	fmt.Println(numTilePossibilities("AAB"))
}
