package _0_basic

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

// https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/description/?envType=problem-list-v2&envId=heap-priority-queue&difficulty=EASY

// 使用回溯  时间太长
func maxSubsequence01(nums []int, k int) []int {
	var backtracking func(nums []int, startIdx, sum int)

	maxSum := math.MinInt32
	var res []int
	var path []int
	backtracking = func(nums []int, startIdx, sum int) {
		if len(path) == k {
			if sum > maxSum {
				maxSum = sum
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = tmp
			}
			return
		}

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(nums, i+1, sum+nums[i])
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0, 0)
	return res
}

// 只需要找到前k大的数字，加起来就是最大的
func maxSubsequence(nums []int, k int) []int {
	type peer struct {
		val int
		idx int
	}
	ps := make([]peer, len(nums))
	for idx, num := range nums {
		ps[idx] = peer{idx: idx, val: num}
	}

	slices.SortFunc(ps, func(a, b peer) int {
		if a.val == b.val {
			if a.idx > b.idx {
				return 1
			} else {
				return -1
			}
		}
		if a.val > b.val {
			return -1
		}
		return 1
	})

	tmp := ps[:k]
	slices.SortFunc(tmp, func(a, b peer) int {
		if a.idx > b.idx {
			return 1
		}
		if a.idx == b.idx {
			return 0
		}
		return -1
	})

	res := make([]int, k)
	for idx, p := range tmp {
		res[idx] = p.val
	}

	return res
}

func TestMaxSubsequence(t *testing.T) {
	fmt.Println(maxSubsequence([]int{-1, -2, 3, 4}, 3))
}
