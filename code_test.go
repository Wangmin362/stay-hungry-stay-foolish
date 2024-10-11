package main

import (
	"testing"
)

// 题目分析：背包的容量为s, 完全背包问题， 没有顺序
// 明确定义：dp[j]表示前i个物品可以凑出来s[:j]字符串
// 递推公式：if dp[i]==true && s[i:j] in wordDict ，那么dp[i] = true
// 初始化: dp[0]= true
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for j := 1; j <= len(s); j++ {
		for i := 0; i <= j; i++ {

		}
	}

	return dp[len(s)]
}
func TestCode(t *testing.T) {
	// head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
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
