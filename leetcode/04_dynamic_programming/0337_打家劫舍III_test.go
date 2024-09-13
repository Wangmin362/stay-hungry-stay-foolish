package _0_basic

import (
	"testing"
)

// https://leetcode.cn/problems/house-robber-iii/description/

// 直接想象成一颗二叉树，使用后续遍历，讨论中间节点偷还是不偷
func robIII01(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	robRoot := root.Val // 偷根节点
	if root.Left != nil {
		robRoot += robIII01(root.Left.Left) + robIII01(root.Left.Right)
	}
	if root.Right != nil {
		robRoot += robIII01(root.Right.Left) + robIII01(root.Right.Right)
	}

	nonRobRoot := 0 // 不偷根节点
	nonRobRoot += robIII01(root.Left) + robIII01(root.Right)
	return max(robRoot, nonRobRoot)
}

// 优化搜索记忆，因为遍历左右节点时，肯定计算了孙子节点
func robIII02(root *TreeNode) int {
	var rob func(root *TreeNode) int
	memory := make(map[*TreeNode]int)
	rob = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if val, ok := memory[root]; ok {
			return val
		}

		if root.Left == nil && root.Right == nil {
			return root.Val
		}

		robRoot := root.Val // 偷根节点
		if root.Left != nil {
			robRoot += rob(root.Left.Left) + rob(root.Left.Right)
		}
		if root.Right != nil {
			robRoot += rob(root.Right.Left) + rob(root.Right.Right)
		}

		nonRobRoot := 0 // 不偷根节点
		nonRobRoot += rob(root.Left) + rob(root.Right)
		memory[root] = max(robRoot, nonRobRoot)
		return max(robRoot, nonRobRoot)
	}

	return rob(root)
}

// 递推关系：根节点的最大价值屈居于根节点偷还是不偷，根据点不偷，那么最大价值是左右子树最大价值之和，根节点偷，那么就是左右节点的子节点之和
// 明确定义：dp[j]表示以j为根节点的最大价值 dp[j] = max(dp[j.Left] + dp[j.Right], dp[j.Left.Left]+dp[j.Left.Right]+dp[j.Right.Left]+dp[j.Right.Right])
// 遍历顺序：由于当前节点偷与不偷取决于左右节点偷与不偷的价值，因此需要使用后序遍历，即先计算出来叶子节点的价值，在计算根节点的价值
func robIII(root *TreeNode) int {
	// res[0]表示不偷，res[1]表示偷
	var rob func(root *TreeNode) [2]int

	rob = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		if root.Left == nil && root.Right == nil { // 讨论叶子节点偷于不偷的情况
			return [2]int{0, root.Val}
		}

		left := rob(root.Left)
		right := rob(root.Right)
		robRoot := root.Val + left[0] + right[0]                  // 偷了当前根节点，那么一定不能投孩子节点
		nonRob := max(left[0], left[1]) + max(right[0], right[1]) // 若不偷根节点，那么等于左右节点的最大值
		return [2]int{nonRob, robRoot}
	}

	res := rob(root)
	return max(res[0], res[1])
}

func TestRobIII(t *testing.T) {
	var testdata = []struct {
		root *TreeNode
		want int
	}{
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}},
			want: 7},
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}},
			want: 9},
	}

	for _, tt := range testdata {
		get := robIII(tt.root)
		if get != tt.want {
			t.Fatalf("want:%v, get:%v", tt.want, get)
		}
	}
}
