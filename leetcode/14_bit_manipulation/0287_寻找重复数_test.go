package _0_basic

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/find-the-duplicate-number/description/?envType=problem-list-v2&envId=bit-manipulation&difficulty=MEDIUM

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func TestFindDuplicate(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
