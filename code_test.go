package main

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {

	maxAr := 0
	left, right := 0, len(height)-1
	for left < right {
		m := min(height[left], height[right]) * (right - left)
		if m > maxAr {
			maxAr = m
		}
		if height[left] > height[right] { // 移动短板才有可能变大，移动长的一定变小
			right--
		} else {
			left++
		}
	}

	return maxAr
}
func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node N叉树定义
type Node struct {
	Val      int
	Children []*Node
}
