package _0_basic

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/relative-ranks/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

func findRelativeRanks(score []int) []string {
	type rank struct {
		score int
		idx   int
		k     string
	}
	ranks := make([]rank, len(score))
	for idx, s := range score {
		ranks[idx] = rank{idx: idx, score: s}
	}
	slices.SortFunc(ranks, func(a, b rank) int {
		if a.score < b.score {
			return 1
		}
		if a.score == b.score {
			return 0
		}
		return -1
	})
	for idx := range ranks {
		switch idx {
		case 0:
			ranks[idx].k = "Gold Medal"
		case 1:
			ranks[idx].k = "Silver Medal"
		case 2:
			ranks[idx].k = "Bronze Medal"
		default:
			ranks[idx].k = strconv.Itoa(idx + 1)
		}
	}
	slices.SortFunc(ranks, func(a, b rank) int {
		if a.idx > b.idx {
			return 1
		}
		if a.score == b.score {
			return 0
		}
		return -1
	})
	res := make([]string, len(score))
	for idx, r := range ranks {
		res[idx] = r.k
	}
	return res
}
func TestFindRelativeRanks(t *testing.T) {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
