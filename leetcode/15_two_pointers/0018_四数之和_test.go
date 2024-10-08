package _0_basic

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)

	var res [][]int
	for one := 0; one+3 < len(nums); one++ {
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}
		for two := one + 1; two+2 < len(nums); two++ {
			if two > one+1 && nums[two] == nums[two-1] {
				continue
			}
			three, four := two+1, len(nums)-1
			for three < four {
				sum := nums[one] + nums[two] + nums[three] + nums[four]
				if sum > target {
					four--
				} else if sum < target {
					three++
				} else {
					res = append(res, []int{nums[one], nums[two], nums[three], nums[four]})
					for three < four && nums[four] == nums[four-1] {
						four--
					}
					for three < four && nums[three] == nums[three+1] {
						three++
					}

					four--
					three++
				}

			}
		}
	}

	return res
}

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
