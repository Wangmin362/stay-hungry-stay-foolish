package main

import (
	"testing"
)

func canJump(nums []int) bool {
	maxStep := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxStep {
			return false
		}
		maxStep = max(maxStep, i+nums[i])
		if maxStep > len(nums)-1 {
			return true
		}
	}

	return true
}

func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
}
