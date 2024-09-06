package main

import (
	"fmt"
)

// https://www.nowcoder.com/practice/54404a78aec1435a81150f15f899417d

type Node struct {
	val  int
	next *Node
}

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		nums := make([]int, 0, n)
		for i := 0; i < n; i++ {
			var num int
			fmt.Scan(&num)
			nums = append(nums, num)
		}

		var k int
		fmt.Scan(&k)

		node := genLinkList(nums)
		knode := getReverseKNode(node, k)
		fmt.Println(knode.val)
	}
}

func genLinkList(nums []int) *Node {
	dummy := &Node{}
	prev := dummy
	for _, num := range nums {
		curr := &Node{val: num}
		prev.next = curr
		prev = curr
	}
	return dummy.next
}

// 1 -> 2 -> 3 -> 4
func getReverseKNode(head *Node, k int) *Node {
	getLen := func(node *Node) int {
		var res int
		for node != nil {
			res++
			node = node.next
		}
		return res
	}
	length := getLen(head)
	if k > length {
		return nil
	}
	step := length - k
	for i := 0; i < step; i++ {
		head = head.next
	}

	return head
}
