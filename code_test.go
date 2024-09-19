package main

import (
	"fmt"
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	sp := strings.Split(s, " ")

	return len(sp[len(sp)-1])
}

func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
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
