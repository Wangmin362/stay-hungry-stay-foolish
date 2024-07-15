package _1_array

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var traversal func(node *TreeNode, path []string)

	var res []string
	traversal = func(node *TreeNode, path []string) {
		if node == nil {
			res = append(res, strings.Join(path, "->"))
			return
		}

		if node.Left != nil {
			path = append(path, strconv.Itoa(node.Left.Val))
			traversal(node.Left, path)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			path = append(path, strconv.Itoa(node.Right.Val))
			traversal(node.Right, path)
			path = path[:len(path)-1]
		}
	}
	traversal(root, []string{strconv.Itoa(root.Val)})
	return res
}
