package _1_array

// https://leetcode.cn/problems/delete-node-in-a-bst/description/

// 递归
// 1、要删除的节点就是叶子节点，即左为空，右也为空的情况
// 2、要删除的节点左不为空有为空
// 3、要删除的节点左为空右不为空
// 4、要删除的节点做不为空右也不为空  让右子树继位，左子树成为右子树的子树
func deleteNode(root *TreeNode, key int) *TreeNode {
	var traversal func(node *TreeNode, key int) *TreeNode // 返回值为新的根节点

	traversal = func(node *TreeNode, key int) *TreeNode {
		if node == nil { // 说明没有找到要删除的这个节点
			return nil
		}
		if node.Val == key {
			if node.Left == nil && node.Right == nil { // 情况一，左为空右为空，此时只需要返回Nil即可
				return nil
			} else if node.Right == nil { // 情况二，左不为空，有为空的情况
				return node.Left
			} else if node.Left == nil { // 情况三，左为空，右不为空
				return node.Right
			} else { // 情况四：左不空，右不空的情况
				// 让右子树继位，左子树成为左子树的子树
				curr := node.Right
				for curr.Left != nil {
					curr = curr.Left
				}
				curr.Left = node.Left // 让左子树成为右子树的子树
				return node.Right
			}
		}

		if node.Val > key {
			node.Left = traversal(node.Left, key)
		} else {
			node.Right = traversal(node.Right, key)
		}
		return node
	}

	return traversal(root, key)
}

// 递归
// 1、要删除的节点就是叶子节点，即左为空，右也为空的情况
// 2、要删除的节点左不为空有为空
// 3、要删除的节点左为空右不为空
// 4、要删除的节点做不为空右也不为空  让左子树继位，右子树成为左子树的子树
func deleteNode02(root *TreeNode, key int) *TreeNode {
	var traversal func(node *TreeNode, key int) *TreeNode // 返回值为新的根节点

	traversal = func(node *TreeNode, key int) *TreeNode {
		if node == nil { // 说明没有找到要删除的这个节点
			return nil
		}
		if node.Val == key {
			if node.Left == nil && node.Right == nil { // 情况一，左为空右为空，此时只需要返回Nil即可
				return nil
			} else if node.Right == nil { // 情况二，左不为空，有为空的情况
				return node.Left
			} else if node.Left == nil { // 情况三，左为空，右不为空
				return node.Right
			} else { // 情况四：左不空，右不空的情况
				// 让左子树继位，右子树成为左子树的子树
				curr := node.Left
				for curr.Right != nil {
					curr = curr.Right
				}
				curr.Right = node.Right // 让左子树成为右子树的子树
				return node.Left
			}
		}

		if node.Val > key {
			node.Left = traversal(node.Left, key)
		} else {
			node.Right = traversal(node.Right, key)
		}
		return node
	}

	return traversal(root, key)
}
