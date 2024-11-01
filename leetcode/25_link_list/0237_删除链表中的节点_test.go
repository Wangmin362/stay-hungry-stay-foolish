package _1_array

// 解题思路：很简单，其实就是先复制下一个节点的值，然后删除一个节点即可
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
