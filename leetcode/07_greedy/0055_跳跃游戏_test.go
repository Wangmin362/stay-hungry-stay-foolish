package _0_basic

import "testing"

func canJump(nums []int) bool {
	maxStep := nums[0] // 能走到的最大位置，位置从0开始

	for i := 1; i < len(nums); i++ {
		if maxStep < i {
			return false
		}
		if maxStep >= len(nums)-1 {
			return true
		}

		maxStep = max(maxStep, i+nums[i])
	}

	return true
}

func TestCanJump(t *testing.T) {

}
