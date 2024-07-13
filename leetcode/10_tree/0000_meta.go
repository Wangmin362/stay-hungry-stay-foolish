package _1_array

import (
	"container/list"
	"strconv"
	"strings"
)

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

func MakeTreeFromArray(arrStr string) *TreeNode {
	arr := append([]string{}, strings.Split(arrStr, ",")...)

	if len(arr) == 0 || arr[0] == "null" {
		return nil
	}

	queue := list.New()
	atoi, _ := strconv.Atoi(arr[0])
	tree := &TreeNode{Val: atoi}
	queue.PushBack(tree)
	queue.PushBack(&TreeNode{Val: 0})
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		idx := queue.Remove(queue.Front()).(*TreeNode).Val

		lidx := 2*idx + 1
		if lidx < len(arr) && arr[lidx] != "null" {
			atoi, _ = strconv.Atoi(arr[lidx])
			node.Left = &TreeNode{Val: atoi}
			queue.PushBack(node.Left)
			queue.PushBack(&TreeNode{Val: lidx})
		}
		ridx := 2*idx + 2
		if lidx+1 < len(arr) && arr[ridx] != "null" {
			atoi, _ = strconv.Atoi(arr[ridx])
			node.Right = &TreeNode{Val: atoi}
			queue.PushBack(node.Right)
			queue.PushBack(&TreeNode{Val: ridx})
		}
	}

	return tree
}
