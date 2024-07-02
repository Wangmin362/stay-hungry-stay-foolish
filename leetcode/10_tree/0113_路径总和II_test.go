package _1_array

// 地址：https://leetcode.cn/problems/path-sum-ii/

// 前序遍历
func pathSum01(root *TreeNode, targetSum int) [][]int {
	// golang中的slice时值拷贝，只要扩容了，底层数组就会发生变化，因此这里不能使用切片作为参数
	//var tarversal func(node *TreeNode, []int, sum int)

	var traversal func(node *TreeNode, sum int)
	traversal = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {

		}
	}
}

func pathSum02(root *TreeNode, targetSum int) [][]int {

}
