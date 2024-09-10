package _1_array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/?envType=problem-list-v2&envId=backtracking&difficulty=MEDIUM
func canPartitionKSubsets(nums []int, k int) bool {
	var backtracking func(ki int)

	var res bool
	cache := make([]bool, len(nums))
	cacheValid := func() bool {
		for i := 0; i < len(cache); i++ {
			if !cache[i] {
				return false
			}
		}
		return true
	}
	path := make([]int, k)
	backtracking = func(ki int) {
		if res {
			return
		}

		if cacheValid() {
			valid := true
			for i := 1; i < k; i++ {
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

		for _, choice := range []bool{true, false} {
			for i := 0; i < len(nums); i++ {
				if cache[i] { // 用过的数字不能再次使用
					continue
				}

				if choice {
					cache[i] = true
					path[ki%k] += nums[i]
					backtracking(ki + 1)
					path[ki%k] -= nums[i]
					cache[i] = false
				} else {
					backtracking(ki + 1)
				}

			}
		}

	}

	backtracking(0)
	return res
}

func TestCanPartitionKSubsets(t *testing.T) {
	fmt.Println(canPartitionKSubsets([]int{19, 11, 1, 3, 3, 1, 9, 9, 3, 1}, 3))
}
