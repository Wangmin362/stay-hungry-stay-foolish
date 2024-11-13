package _1_array

type Node3294 struct {
	Val  int
	Next *Node3294
	Prev *Node3294
}

func toArray3294(head *Node3294) []int {
	var res []int

	h1 := head.Prev
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for h1 != nil {
		res = append([]int{h1.Val}, res...)
		h1 = h1.Prev
	}

	return res
}
