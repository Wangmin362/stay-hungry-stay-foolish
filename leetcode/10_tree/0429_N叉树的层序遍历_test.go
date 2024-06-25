package _1_array

// 地址：https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/
//
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//// 很简单，就是层序遍历
//func levelOrderN(root *Node) [][]int {
//	if root == nil {
//		return nil
//	}
//	var res [][]int
//	queue := list.New()
//	queue.PushBack(root)
//	for queue.Len() > 0 {
//		length := queue.Len()
//		temp := make([]int, length)
//		for i := 0; i < length; i++ {
//			node := queue.Remove(queue.Front()).(*Node)
//			temp[i] = node.Val
//			if len(node.Children) != 0 {
//				for _, nc := range node.Children {
//					if nc != nil {
//						queue.PushBack(nc)
//					}
//				}
//			}
//		}
//		res = append(res, temp)
//	}
//
//	return res
//}
