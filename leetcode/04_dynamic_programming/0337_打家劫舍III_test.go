package _0_basic

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

// 优化搜索记忆，以为遍历左右节点时，肯定计算了孙子节点
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

// 动态规划
func robIII03(root *TreeNode) int {
	var rob func(root *TreeNode) [2]int
	rob = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		leftMax := rob(root.Left)   // 左孩子偷与不透的最大值
		rightMax := rob(root.Right) // 右孩子偷与不偷的最大值

		robRoot := root.Val + leftMax[0] + rightMax[0]
		nonRobRoot := max(leftMax[0], leftMax[1]) + max(rightMax[0], rightMax[1])
		return [2]int{nonRobRoot, robRoot}
	}

	res := rob(root)
	return max(res[0], res[1])
}
